package version

import "fmt"

var version = "latest"

type Printer struct{}

func (v Printer) Name() string {
	return "version"
}

func (v Printer) Run([]string) error {
	fmt.Printf("Running Milo %v\n", version)

	return nil
}

func (v Printer) HelpText() string {
	return "prints the Milo version"
}
