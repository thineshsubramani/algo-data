package helper

import (
	"fmt"
	"strings"
)

// ComponentInfo holds metadata about a specific data structure.
type ComponentInfo struct {
	Name      string
	Functions []string
	Options   []string
}

var registry = make(map[string]ComponentInfo)

// Register allows a package to announce its capabilities to the helper system.
func Register(info ComponentInfo) {
	registry[strings.ToLower(info.Name)] = info
}

// Describe prints the available functions and options for a specific component.
func Describe(name string) {
	info, ok := registry[strings.ToLower(name)]
	if !ok {
		fmt.Printf("No help information found for: %s\n", name)
		return
	}

	fmt.Printf("\n--- HELP: %s ---\n", strings.ToUpper(info.Name))
	fmt.Println("Functions:")
	for _, f := range info.Functions {
		fmt.Printf("  • %s\n", f)
	}
	if len(info.Options) > 0 {
		fmt.Println("Options:")
		for _, o := range info.Options {
			fmt.Printf("  • %s\n", o)
		}
	}
	fmt.Println("------------------")
}