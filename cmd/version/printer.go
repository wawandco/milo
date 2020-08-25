package version

import "fmt"

// version holds the version of the tool.
var version = "latest"

// Printer prints the version of the tool in version, this one will typically be
// overridden at buildtime with ldflags.
type Printer struct{}

// Name of the command for the CLI (milo version).
func (v Printer) Name() string {
	return "version"
}

// Run and print the version prefixed with Running milo.
func (v Printer) Run([]string) error {
	fmt.Printf("Running Milo %v\n", version)

	return nil
}

// HelpText for this runner.
func (v Printer) HelpText() string {
	return "prints the Milo version"
}
