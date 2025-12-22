// Package core provides the configuration types for the installer.
package core

import "gopkg.in/yaml.v3"

// BranchConfig defines conditional branching logic.
type BranchConfig struct {
	Condition string            `yaml:"condition" json:"condition"`                 // Expression to evaluate
	Branches  map[string]string `yaml:"branches" json:"branches"`                   // value -> step ID
	Default   string            `yaml:"default,omitempty" json:"default,omitempty"` // Default step if no match
}

// UnmarshalYAML supports both condition/branches/default and legacy when/then/else forms.
func (b *BranchConfig) UnmarshalYAML(value *yaml.Node) error {
	var raw struct {
		Condition string            `yaml:"condition"`
		Branches  map[string]string `yaml:"branches"`
		Default   string            `yaml:"default"`
		When      string            `yaml:"when"`
		Then      string            `yaml:"then"`
		Else      string            `yaml:"else"`
	}
	if err := value.Decode(&raw); err != nil {
		return err
	}

	if raw.Condition != "" || len(raw.Branches) > 0 || raw.Default != "" {
		b.Condition = raw.Condition
		b.Branches = raw.Branches
		b.Default = raw.Default
		return nil
	}

	if raw.When == "" && raw.Then == "" && raw.Else == "" {
		return nil
	}

	b.Condition = raw.When
	if raw.Then != "" {
		b.Branches = map[string]string{"true": raw.Then}
	}
	b.Default = raw.Else
	return nil
}

// TaskConfig represents a task configuration from YAML.
type TaskConfig struct {
	Type          string         `yaml:"type" json:"type"`
	ID            string         `yaml:"id,omitempty" json:"id,omitempty"`
	Params        map[string]any `yaml:",inline" json:",inline"`
	FailurePolicy string         `yaml:"on_failure,omitempty" json:"on_failure,omitempty"`
	Retries       int            `yaml:"retries,omitempty" json:"retries,omitempty"`
}

// StepConfig represents a step configuration.
type StepConfig struct {
	ID        string           `yaml:"id" json:"id"`
	Title     string           `yaml:"title" json:"title"`
	Screen    *ScreenConfig    `yaml:"screen,omitempty" json:"screen,omitempty"`
	Tasks     []TaskConfig     `yaml:"tasks,omitempty" json:"tasks,omitempty"`
	Guards    []map[string]any `yaml:"guards,omitempty" json:"guards,omitempty"`
	Next      string           `yaml:"next,omitempty" json:"next,omitempty"`
	Prev      string           `yaml:"prev,omitempty" json:"prev,omitempty"`
	Branch    *BranchConfig    `yaml:"branch,omitempty" json:"branch,omitempty"`
	AllowJump bool             `yaml:"allowJump,omitempty" json:"allowJump,omitempty"`
	Route     string           `yaml:"route,omitempty" json:"route,omitempty"`
}

// ScreenConfig represents screen configuration.
type ScreenConfig struct {
	Type               string        `yaml:"type" json:"type"`
	Title              string        `yaml:"title,omitempty" json:"title,omitempty"`
	Description        string        `yaml:"description,omitempty" json:"description,omitempty"`
	Content            string        `yaml:"content,omitempty" json:"content,omitempty"`
	ContentFile        string        `yaml:"contentFile,omitempty" json:"contentFile,omitempty"`
	Source             string        `yaml:"source,omitempty" json:"source,omitempty"`
	BannerPath         string        `yaml:"bannerPath,omitempty" json:"bannerPath,omitempty"`
	RequireScrollToEnd bool          `yaml:"requireScrollToEnd,omitempty" json:"requireScrollToEnd,omitempty"`
	Bind               string        `yaml:"bind,omitempty" json:"bind,omitempty"`
	Options            []Option      `yaml:"options,omitempty" json:"options,omitempty"`
	Fields             []FieldConfig `yaml:"fields,omitempty" json:"fields,omitempty"`
}

// FieldConfig represents a form field configuration.
type FieldConfig struct {
	Type       string   `yaml:"type" json:"type"`
	Label      string   `yaml:"label" json:"label"`
	Variable   string   `yaml:"variable" json:"variable"`
	Default    string   `yaml:"default,omitempty" json:"default,omitempty"`
	Required   bool     `yaml:"required,omitempty" json:"required,omitempty"`
	Hint       string   `yaml:"hint,omitempty" json:"hint,omitempty"`
	Validation string   `yaml:"validation,omitempty" json:"validation,omitempty"`
	Options    []Option `yaml:"options,omitempty" json:"options,omitempty"`
}

// Option represents an option in options screen or select field.
type Option struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Default bool   `yaml:"default,omitempty" json:"default,omitempty"`
}

// FlowConfig represents a workflow configuration.
type FlowConfig struct {
	Entry string        `yaml:"entry" json:"entry"`
	Steps []*StepConfig `yaml:"steps" json:"steps"`
}

// ProductConfig represents product metadata.
type ProductConfig struct {
	Name         string       `yaml:"name" json:"name"`
	Version      string       `yaml:"version,omitempty" json:"version,omitempty"`
	Vendor       string       `yaml:"vendor,omitempty" json:"vendor,omitempty"`
	Homepage     string       `yaml:"homepage,omitempty" json:"homepage,omitempty"`
	License      string       `yaml:"license,omitempty" json:"license,omitempty"`
	Description  string       `yaml:"description,omitempty" json:"description,omitempty"`
	Icon         string       `yaml:"icon,omitempty" json:"icon,omitempty"`
	Logo         string       `yaml:"logo,omitempty" json:"logo,omitempty"`
	MinOSVersion string       `yaml:"minOSVersion,omitempty" json:"minOSVersion,omitempty"`
	Theme        *ThemeConfig `yaml:"theme,omitempty" json:"theme,omitempty"`
}

// ThemeConfig contains theming options.
type ThemeConfig struct {
	PrimaryColor string `yaml:"primaryColor,omitempty" json:"primaryColor,omitempty"`
}

// SourcesConfig represents download sources configuration.
type SourcesConfig struct {
	BaseURL    string            `yaml:"baseUrl,omitempty" json:"baseUrl,omitempty"`
	Mirrors    []string          `yaml:"mirrors,omitempty" json:"mirrors,omitempty"`
	Components []ComponentConfig `yaml:"components,omitempty" json:"components,omitempty"`
}

// ComponentConfig represents a downloadable component.
type ComponentConfig struct {
	Name        string `yaml:"name" json:"name"`
	Path        string `yaml:"path" json:"path"`
	SHA256      string `yaml:"sha256,omitempty" json:"sha256,omitempty"`
	Size        int64  `yaml:"size,omitempty" json:"size,omitempty"`
	Required    bool   `yaml:"required,omitempty" json:"required,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// InstallConfig represents installation actions configuration.
type InstallConfig struct {
	PreInstall  []TaskConfig `yaml:"preInstall,omitempty" json:"preInstall,omitempty"`
	Install     []TaskConfig `yaml:"install,omitempty" json:"install,omitempty"`
	PostInstall []TaskConfig `yaml:"postInstall,omitempty" json:"postInstall,omitempty"`
}

// UninstallConfig represents uninstallation configuration.
type UninstallConfig struct {
	Tasks []TaskConfig `yaml:"tasks,omitempty" json:"tasks,omitempty"`
}

// Config represents the complete installer configuration.
type Config struct {
	Schema    string                 `yaml:"$schema,omitempty" json:"$schema,omitempty"`
	Product   *ProductConfig         `yaml:"product" json:"product"`
	Meta      map[string]any         `yaml:"meta,omitempty" json:"meta,omitempty"`
	Sources   *SourcesConfig         `yaml:"sources,omitempty" json:"sources,omitempty"`
	Flow      *FlowConfig            `yaml:"flow,omitempty" json:"flow,omitempty"`
	Flows     map[string]*FlowConfig `yaml:"flows,omitempty" json:"flows,omitempty"`
	Install   *InstallConfig         `yaml:"install,omitempty" json:"install,omitempty"`
	Uninstall *UninstallConfig       `yaml:"uninstall,omitempty" json:"uninstall,omitempty"`
}
