// Package core provides environment detection helpers.
package core

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// DetectEnv fills InstallContext.Env with detected values.
func DetectEnv(ctx *InstallContext) {
	ctx.Env.Arch = runtime.GOARCH
	ctx.Env.IsRoot = os.Geteuid() == 0
	ctx.Env.HasSudo = hasExecutable("sudo")
	ctx.Env.HasPolkit = hasExecutable("pkexec")
	ctx.Env.Desktop = detectDesktop()

	distro, version := detectDistro()
	ctx.Env.Distro = distro
	ctx.Env.DistroVersion = version

	ctx.Env.DiskFreeMB = detectDiskFreeMB()
}

func hasExecutable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func detectDesktop() string {
	if val := os.Getenv("XDG_CURRENT_DESKTOP"); val != "" {
		return strings.ToLower(val)
	}
	if val := os.Getenv("DESKTOP_SESSION"); val != "" {
		return strings.ToLower(val)
	}
	return "unknown"
}

func detectDistro() (string, string) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "unknown", ""
	}
	defer file.Close()

	var id string
	var version string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			id = strings.Trim(line[3:], "\"")
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			version = strings.Trim(line[len("VERSION_ID="):], "\"")
		}
	}

	if id == "" {
		id = "unknown"
	}
	return id, version
}

func detectDiskFreeMB() int64 {
	path, err := os.UserHomeDir()
	if err != nil || path == "" {
		path = "/"
	}

	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0
	}

	available := stat.Bavail * uint64(stat.Bsize)
	return int64(available / (1024 * 1024))
}
