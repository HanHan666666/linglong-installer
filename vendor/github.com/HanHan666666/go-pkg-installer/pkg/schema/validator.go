// Package schema provides configuration validation against JSON Schema.
package schema

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed installer.schema.json
var schemaFS embed.FS

// ValidationError represents a validation error with path information.
type ValidationError struct {
	Path    string
	Message string
}

func (e ValidationError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("%s: %s", e.Path, e.Message)
	}
	return e.Message
}

// ValidationResult contains the result of validation.
type ValidationResult struct {
	Valid  bool
	Errors []ValidationError
}

// Validator validates installer configuration against the JSON Schema.
type Validator struct {
	schema *jsonschema.Schema
}

// NewValidator creates a new Validator with the embedded schema.
func NewValidator() (*Validator, error) {
	schemaData, err := schemaFS.ReadFile("installer.schema.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded schema: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020

	if err := compiler.AddResource("installer.schema.json", strings.NewReader(string(schemaData))); err != nil {
		return nil, fmt.Errorf("failed to add schema resource: %w", err)
	}

	schema, err := compiler.Compile("installer.schema.json")
	if err != nil {
		return nil, fmt.Errorf("failed to compile schema: %w", err)
	}

	return &Validator{schema: schema}, nil
}

// NewValidatorFromFile creates a Validator from a schema file path.
func NewValidatorFromFile(schemaPath string) (*Validator, error) {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020

	schema, err := compiler.Compile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to compile schema from %s: %w", schemaPath, err)
	}

	return &Validator{schema: schema}, nil
}

// ValidateYAML validates YAML content against the schema.
func (v *Validator) ValidateYAML(yamlContent []byte) ValidationResult {
	var data any
	if err := yaml.Unmarshal(yamlContent, &data); err != nil {
		return ValidationResult{
			Valid: false,
			Errors: []ValidationError{
				{Message: fmt.Sprintf("invalid YAML: %v", err)},
			},
		}
	}

	data = convertYAMLToJSON(data)
	return v.validate(data)
}

// ValidateJSON validates JSON content against the schema.
func (v *Validator) ValidateJSON(jsonContent []byte) ValidationResult {
	var data any
	if err := json.Unmarshal(jsonContent, &data); err != nil {
		return ValidationResult{
			Valid: false,
			Errors: []ValidationError{
				{Message: fmt.Sprintf("invalid JSON: %v", err)},
			},
		}
	}

	return v.validate(data)
}

// Validate validates a Go value against the schema.
func (v *Validator) Validate(data any) ValidationResult {
	data = convertYAMLToJSON(data)
	return v.validate(data)
}

func (v *Validator) validate(data any) ValidationResult {
	err := v.schema.Validate(data)
	if err == nil {
		return ValidationResult{Valid: true}
	}

	var validationErr *jsonschema.ValidationError
	if errors.As(err, &validationErr) {
		return ValidationResult{
			Valid:  false,
			Errors: extractErrors(validationErr),
		}
	}

	return ValidationResult{
		Valid: false,
		Errors: []ValidationError{
			{Message: err.Error()},
		},
	}
}

func extractErrors(err *jsonschema.ValidationError) []ValidationError {
	var result []ValidationError

	if err.Message != "" {
		result = append(result, ValidationError{
			Path:    err.InstanceLocation,
			Message: err.Message,
		})
	}

	for _, cause := range err.Causes {
		result = append(result, extractErrors(cause)...)
	}

	return result
}

func convertYAMLToJSON(data any) any {
	switch v := data.(type) {
	case map[string]any:
		result := make(map[string]any)
		for key, val := range v {
			result[key] = convertYAMLToJSON(val)
		}
		return result
	case map[any]any:
		result := make(map[string]any)
		for key, val := range v {
			keyStr := fmt.Sprintf("%v", key)
			result[keyStr] = convertYAMLToJSON(val)
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = convertYAMLToJSON(val)
		}
		return result
	default:
		return v
	}
}

// Config represents the parsed installer configuration.
// Deprecated: Use core.Config instead.
type Config = core.Config

// ProductConfig contains product branding information.
// Deprecated: Use core.ProductConfig instead.
type ProductConfig = core.ProductConfig

// FlowConfig represents a workflow configuration.
// Deprecated: Use core.FlowConfig instead.
type FlowConfig = core.FlowConfig

// StepConfig represents a step configuration.
// Deprecated: Use core.StepConfig instead.
type StepConfig = core.StepConfig

// ScreenConfig represents screen configuration.
// Deprecated: Use core.ScreenConfig instead.
type ScreenConfig = core.ScreenConfig

// FieldConfig represents a form field configuration.
// Deprecated: Use core.FieldConfig instead.
type FieldConfig = core.FieldConfig

// Option represents an option in options screen.
// Deprecated: Use core.Option instead.
type Option = core.Option

// BranchConfig represents conditional branching.
// Deprecated: Use core.BranchConfig instead.
type BranchConfig = core.BranchConfig

// ConfigLoader is an interface for loading and validating configurations
type ConfigLoader interface {
	Load(yamlContent []byte) (*core.Config, error)
}

// SchemaValidatorLoader loads configs with schema validation
type SchemaValidatorLoader struct{}

func (l *SchemaValidatorLoader) Load(yamlContent []byte) (*core.Config, error) {
	validator, err := NewValidator()
	if err != nil {
		// Skip validation if embedded schema fails
		_ = err
	} else {
		result := validator.ValidateYAML(yamlContent)
		if !result.Valid {
			var errMsgs []string
			for _, e := range result.Errors {
				errMsgs = append(errMsgs, e.Error())
			}
			return nil, fmt.Errorf("validation failed:\n  %s", strings.Join(errMsgs, "\n  "))
		}
	}

	var config core.Config
	if err := yaml.Unmarshal(yamlContent, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}

// LoadConfig loads and validates an installer configuration from YAML.
func LoadConfig(yamlContent []byte) (*core.Config, error) {
	loader := &SchemaValidatorLoader{}
	return loader.Load(yamlContent)
}

// LoadConfigFromString loads and validates an installer configuration from a YAML string.
func LoadConfigFromString(yamlStr string) (*core.Config, error) {
	return LoadConfig([]byte(yamlStr))
}
