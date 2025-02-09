package tygoja

const (
	defaultIndent = "  "

	// custom base types that every package has access to
	BaseTypeDict = "_TygojaDict" // Record type alternative as a more generic map-like type
	BaseTypeAny  = "_TygojaAny"  // any type alias to allow easier extends generation
)

// FieldNameFormatterFunc defines a function for formatting a field name.
type FieldNameFormatterFunc func(string) string

// MethodNameFormatterFunc defines a function for formatting a method name.
type MethodNameFormatterFunc func(string) string

type Config struct {
	// Packages is a list of package paths just like you would import them in Go.
	// Use "*" to generate all package types.
	//
	// Example:
	//
	// 	Packages: map[string][]string{
	// 		"time": {"Time"},
	// 		"github.com/hanzoai/base/core": {"*"},
	// 	}
	Packages map[string][]string

	// Heading specifies a content that will be put at the top of the output declaration file.
	//
	// You would generally use this to import custom types or some custom TS declarations.
	Heading string

	// TypeMappings specifies custom type translations.
	//
	// Useful for for mapping 3rd party package types, eg "unsafe.Pointer" => "CustomType".
	//
	// Be default unrecognized types will be recursively generated by
	// traversing their import package (when possible).
	TypeMappings map[string]string

	// WithConstants indicates whether to generate types for constants
	// ("false" by default).
	WithConstants bool

	// WithPackageFunctions indicates whether to generate types
	// for package level functions ("false" by default).
	WithPackageFunctions bool

	// FieldNameFormatter allows specifying a custom struct field name formatter.
	FieldNameFormatter FieldNameFormatterFunc

	// MethodNameFormatter allows specifying a custom method name formatter.
	MethodNameFormatter MethodNameFormatterFunc

	// StartModifier usually should be "export" or declare but as of now prevents
	// the LSP autocompletion so we keep it empty.
	//
	// See also:
	// https://github.com/microsoft/TypeScript/issues/54330
	// https://github.com/microsoft/TypeScript/pull/49644
	StartModifier string

	// Indent allow customizing the default indentation (use \t if you want tabs).
	Indent string
}

// Initializes the defaults (if not already) of the current config.
func (c *Config) InitDefaults() {
	if c.Indent == "" {
		c.Indent = defaultIndent
	}

	if c.TypeMappings == nil {
		c.TypeMappings = make(map[string]string)
	}

	// special case for the unsafe package because it doesn't return its types in pkg.Syntax
	if _, ok := c.TypeMappings["unsafe.Pointer"]; !ok {
		c.TypeMappings["unsafe.Pointer"] = "number"
	}
}
