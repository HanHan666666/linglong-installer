// Package main provides the reference installer application
// demonstrating the go-pkg-installer framework capabilities.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/builtin"
	"github.com/HanHan666666/go-pkg-installer/pkg/core"
	"github.com/HanHan666666/go-pkg-installer/pkg/schema"
	"github.com/HanHan666666/go-pkg-installer/pkg/ui"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	// Parse command line flags
	configPath := flag.String("config", "", "Path to installer configuration YAML file")
	action := flag.String("action", "install", "Action to perform: install, uninstall")
	validateOnly := flag.Bool("validate", false, "Only validate the configuration file")
	showVersion := flag.Bool("version", false, "Show version information")
	headless := flag.Bool("headless", false, "Run in headless/CLI mode (no GUI)")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	acceptLicense := flag.Bool("accept-license", false, "Accept license agreement (CLI)")
	installDir := flag.String("install-dir", "", "Installation directory (CLI)")
	installType := flag.String("install-type", "", "Installation type (CLI)")
	privilege := flag.String("privilege", "", "Privilege strategy: sudo|pkexec|none")
	var overrides kvFlags
	flag.Var(&overrides, "set", "Set context value (key=value), repeatable")
	flag.Parse()

	// Show version
	if *showVersion {
		fmt.Printf("go-pkg-installer %s (built %s)\n", version, buildTime)
		os.Exit(0)
	}

	// Load configuration
	cfg, err := loadConfig(*configPath, embeddedConfig, *verbose)
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	if *validateOnly {
		fmt.Println("✓ Configuration is valid")
		os.Exit(0)
	}

	// Register builtin tasks
	builtin.RegisterAll()
	RegisterDistroScriptTask()
	// Register builtin guards
	core.RegisterBuiltinGuards()

	// Create installation context
	ctx := core.NewInstallContext()
	defer ctx.CloseLogFile()

	// Set product information
	if cfg.Product != nil {
		ctx.Set("product.name", cfg.Product.Name)

		// Handle logo - use embedded logo data
		logoPath := cfg.Product.Logo
		if logoPath == "assets/logo.png" || logoPath == "" {
			ctx.Set("product.logo.bytes", embeddedLogo)
		} else {
			ctx.Set("product.logo", logoPath)
		}

		if cfg.Product.Theme != nil && cfg.Product.Theme.PrimaryColor != "" {
			ctx.Set("theme.primaryColor", cfg.Product.Theme.PrimaryColor)
		}
	}

	// Set meta
	for key, value := range cfg.Meta {
		ctx.Set(key, value)
		ctx.Set("meta."+key, value)
	}

	// CLI overrides
	if *privilege != "" {
		ctx.Set("privilege.strategy", *privilege)
	}
	if *acceptLicense {
		ctx.Set("license.accepted", true)
	}
	if *installDir != "" {
		ctx.Set("install.dir", *installDir)
		ctx.Set("install_dir", *installDir)
	}
	if *installType != "" {
		ctx.Set("install.type", *installType)
	}
	for _, kv := range overrides {
		key, value := parseOverride(kv)
		if key != "" {
			ctx.Set(key, value)
		}
	}

	// Preflight environment detection
	core.DetectEnv(ctx)
	runDistroDetection(ctx, *verbose)
	ensureEmbeddedScript(ctx, "scripts/user/setup-linyaps-dbus.sh", "linyaps.dbus.setup_script", *verbose)
	ensureEmbeddedScript(ctx, "scripts/user/install-linyaps.sh", "linyaps.store.setup_script", *verbose)

	// Create event bus
	eventBus := core.NewEventBus()
	ctx.SetEventBus(eventBus)

	// Create workflow
	workflow := core.NewWorkflow(ctx, eventBus)

	// Add flows from config
	for flowName, flowCfg := range cfg.Flows {
		flow := &core.Flow{
			ID:    flowName,
			Entry: flowCfg.Entry,
		}

		for _, stepCfg := range flowCfg.Steps {
			step := &core.Step{
				ID:     stepCfg.ID,
				Title:  ctx.Render(stepCfg.Title),
				Config: stepCfg,
			}

			// Copy guards config
			for _, g := range stepCfg.Guards {
				step.GuardsCfg = append(step.GuardsCfg, g)
			}

			// Copy tasks config
			for _, t := range stepCfg.Tasks {
				taskMap := map[string]any{
					"type": t.Type,
					"id":   t.ID,
				}
				for k, v := range t.Params {
					taskMap[k] = v
				}
				step.TasksCfg = append(step.TasksCfg, taskMap)
			}

			flow.Steps = append(flow.Steps, step)
		}

		if err := workflow.AddFlow(flow); err != nil {
			log.Fatalf("Failed to add flow %s: %v", flowName, err)
		}
	}

	// Select the requested flow
	if err := workflow.SelectFlow(*action); err != nil {
		log.Fatalf("Failed to select flow '%s': %v", *action, err)
	}
	ctx.Runtime.Action = *action
	ctx.Plan = core.BuildTaskPlan(cfg.Flows[*action])

	// Setup log file output
	if logPath := defaultLogPath(cfg); logPath != "" {
		if err := ctx.SetLogFile(logPath); err != nil && *verbose {
			log.Printf("Failed to set log file: %v", err)
		}
	}

	// Elevate if needed
	maybeElevate(ctx, cfg, *action)

	if *verbose {
		log.Printf("Selected flow: %s", *action)
		log.Printf("Steps: %d", len(workflow.Steps()))
	}

	// Subscribe to events for logging
	if *verbose {
		eventBus.Subscribe(core.EventLog, func(e core.Event) {
			if p := e.LogPayload(); p != nil {
				log.Printf("[LOG] %s", p.Message)
			}
		})
		eventBus.Subscribe(core.EventProgress, func(e core.Event) {
			if p := e.ProgressPayload(); p != nil {
				log.Printf("[PROGRESS] %.0f%% - %s", p.Progress*100, p.Message)
			}
		})
		eventBus.Subscribe(core.EventTaskStart, func(e core.Event) {
			if t := e.TaskPayload(); t != nil {
				log.Printf("[TASK] Starting: %s", t.TaskID)
			}
		})
		eventBus.Subscribe(core.EventTaskComplete, func(e core.Event) {
			if t := e.TaskPayload(); t != nil {
				log.Printf("[TASK] Completed: %s", t.TaskID)
			}
		})
		eventBus.Subscribe(core.EventTaskError, func(e core.Event) {
			if t := e.TaskPayload(); t != nil {
				log.Printf("[TASK] Failed: %s - %v", t.TaskID, t.Error)
			}
		})
	}

	// Run in appropriate mode
	if *headless {
		runHeadless(ctx, workflow, eventBus, cfg, *verbose)
	} else {
		runGUI(ctx, workflow, eventBus)
	}
}

// loadConfig loads configuration from embedded or external source
func loadConfig(configPath string, embeddedConfig []byte, verbose bool) (*core.Config, error) {
	if configPath == "" {
		// Use embedded configuration
		if verbose {
			log.Println("Using embedded configuration")
		}
		return loadAndValidateConfig(embeddedConfig)
	}

	// Use external configuration file
	if verbose {
		log.Printf("Loading configuration from: %s", configPath)
	}
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	return loadAndValidateConfig(content)
}

// loadAndValidateConfig loads and validates the configuration from bytes
func loadAndValidateConfig(content []byte) (*core.Config, error) {
	// Use schema.LoadConfig which validates and parses
	cfg, err := schema.LoadConfig(content)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// runGUI runs the installer in GUI mode
func runGUI(ctx *core.InstallContext, workflow *core.Workflow, eventBus *core.EventBus) {
	// Create installer window
	win := ui.NewInstallerWindow(ctx, workflow, eventBus)
	win.RegisterScreenRenderer("linglongVersion", NewLinglongVersionScreen)

	// Set callbacks
	win.OnComplete(func() {
		log.Println("Installation completed successfully")
	})

	win.OnCancel(func() {
		log.Println("Installation cancelled by user")
		os.Exit(1)
	})

	// Run the window
	if err := win.Run(); err != nil {
		log.Fatalf("UI error: %v", err)
	}
}

// runHeadless runs the installer in CLI/headless mode
func runHeadless(ctx *core.InstallContext, workflow *core.Workflow, eventBus *core.EventBus, cfg *core.Config, verbose bool) {
	productName := "Application"
	if cfg.Product != nil {
		productName = cfg.Product.Name
	}

	fmt.Printf("=== %s Installation ===\n", productName)
	fmt.Println()

	// Create task runner
	runner := core.NewTaskRunner(ctx, eventBus)

	// Process each step
	for !workflow.IsComplete() {
		step := workflow.CurrentStep()
		if step == nil {
			log.Fatal("Workflow error: no current step")
		}

		fmt.Printf("Step: %s\n", step.Title)

		// If step has tasks, execute them
		if step.Config != nil && len(step.Config.Tasks) > 0 {
			fmt.Println("  Executing tasks...")

			// Queue tasks from config
			for _, taskCfg := range step.Config.Tasks {
				if err := runner.QueueConfig(taskCfg); err != nil {
					log.Fatalf("Failed to queue task: %v", err)
				}
			}

			// Run tasks
			if err := runner.Run(); err != nil {
				log.Fatalf("Task execution failed: %v", err)
			}

			fmt.Println("  ✓ Tasks completed")
		}

		// Move to next step
		if !workflow.IsLastStep() {
			if _, err := workflow.Next(); err != nil {
				log.Fatalf("Failed to advance: %v", err)
			}
		} else {
			workflow.Complete()
		}
	}

	fmt.Println()
	fmt.Println("=== Installation Complete ===")
}

type kvFlags []string

func (k *kvFlags) String() string {
	return strings.Join(*k, ",")
}

func (k *kvFlags) Set(value string) error {
	*k = append(*k, value)
	return nil
}

func parseOverride(input string) (string, any) {
	parts := strings.SplitN(input, "=", 2)
	if len(parts) != 2 {
		return "", nil
	}
	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	if key == "" {
		return "", nil
	}

	if strings.EqualFold(val, "true") {
		return key, true
	}
	if strings.EqualFold(val, "false") {
		return key, false
	}
	if i, err := strconv.ParseInt(val, 10, 64); err == nil {
		return key, i
	}
	return key, val
}

func maybeElevate(ctx *core.InstallContext, cfg *core.Config, action string) {
	if core.Elevated() || ctx.Env.IsRoot {
		return
	}
	if !core.NeedsPrivilege(cfg, action) {
		return
	}

	execPath := resolveExecutablePath()
	strategy := core.GetPrivilegeStrategy(ctx)
	switch strategy {
	case core.PrivilegeSudo:
		if !ctx.Env.HasSudo {
			log.Fatalf("Privilege required but sudo is not available")
		}
		runElevated("sudo", append([]string{"-E", execPath}, os.Args[1:]...))
	case core.PrivilegePkexec:
		if !ctx.Env.HasPolkit {
			log.Fatalf("Privilege required but pkexec is not available")
		}
		runElevated("pkexec", append([]string{"env", "GPKI_ELEVATED=1", execPath}, os.Args[1:]...))
	case core.PrivilegeNone:
		log.Fatalf("Privilege required but no elevation strategy configured")
	default:
		log.Fatalf("Unknown privilege strategy: %s", strategy)
	}
}

func resolveExecutablePath() string {
	execPath, err := os.Executable()
	if err == nil && execPath != "" {
		return execPath
	}

	if absPath, err := filepath.Abs(os.Args[0]); err == nil && absPath != "" {
		return absPath
	}

	return os.Args[0]
}

func runElevated(cmdName string, args []string) {
	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "GPKI_ELEVATED=1")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Elevation failed: %v", err)
	}
	os.Exit(0)
}

func defaultLogPath(cfg *core.Config) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	productName := "installer"
	if cfg != nil && cfg.Product != nil && cfg.Product.Name != "" {
		productName = cfg.Product.Name
	}
	safeName := strings.ToLower(strings.ReplaceAll(productName, " ", "-"))
	return filepath.Join(home, ".local", "share", "go-pkg-installer", safeName, "installer.log")
}

func runDistroDetection(ctx *core.InstallContext, verbose bool) {
	if ctx == nil {
		return
	}

	factory, ok := core.Tasks.Get("distroScript")
	if !ok {
		if verbose {
			log.Printf("distroScript task not registered; skipping distro detection")
		}
		return
	}

	task, err := factory(map[string]any{}, ctx)
	if err != nil {
		if verbose {
			log.Printf("Failed to create distroScript task: %v", err)
		}
		return
	}

	if err := task.Validate(); err != nil {
		if verbose {
			log.Printf("Invalid distroScript task: %v", err)
		}
		return
	}

	if err := task.Execute(ctx, nil); err != nil && verbose {
		log.Printf("distroScript execution failed: %v", err)
	}
}

func ensureEmbeddedScript(ctx *core.InstallContext, relPath string, ctxKey string, verbose bool) {
	if ctx == nil || relPath == "" || ctxKey == "" {
		return
	}

	needsCommon := strings.HasPrefix(relPath, "scripts/user/")
	if resolved, ok := resolveScriptPath(relPath); ok {
		ctx.Set(ctxKey, resolved)
		return
	}

	data, err := embeddedScripts.ReadFile(relPath)
	if err != nil {
		if verbose {
			log.Printf("Embedded script not found (%s): %v", relPath, err)
		}
		return
	}

	tmpDir, err := os.MkdirTemp("", "linglong-installer-")
	if err != nil {
		if verbose {
			log.Printf("Failed to create temp dir for %s: %v", relPath, err)
		}
		return
	}

	relPath = filepath.FromSlash(relPath)
	destPath := filepath.Join(tmpDir, relPath)
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		if verbose {
			log.Printf("Failed to create temp dir for %s: %v", relPath, err)
		}
		return
	}

	if needsCommon {
		commonData, err := embeddedScripts.ReadFile("scripts/common.sh")
		if err != nil {
			if verbose {
				log.Printf("Embedded script not found (scripts/common.sh): %v", err)
			}
			return
		}
		commonPath := filepath.Join(tmpDir, "scripts", "common.sh")
		if err := os.MkdirAll(filepath.Dir(commonPath), 0o755); err != nil {
			if verbose {
				log.Printf("Failed to create temp dir for scripts/common.sh: %v", err)
			}
			return
		}
		if err := os.WriteFile(commonPath, commonData, 0o644); err != nil {
			if verbose {
				log.Printf("Failed to write embedded script scripts/common.sh: %v", err)
			}
			return
		}
	}

	if err := os.WriteFile(destPath, data, 0o755); err != nil {
		if verbose {
			log.Printf("Failed to write embedded script %s: %v", relPath, err)
		}
		return
	}

	ctx.Set(ctxKey, destPath)
}

func resolveScriptPath(relPath string) (string, bool) {
	if relPath == "" {
		return "", false
	}

	candidates := []string{}
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(cwd, relPath))
	}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), relPath))
	}

	for _, candidate := range candidates {
		info, err := os.Stat(candidate)
		if err == nil && !info.IsDir() {
			if absPath, err := filepath.Abs(candidate); err == nil {
				return absPath, true
			}
			return candidate, true
		}
	}

	return "", false
}
