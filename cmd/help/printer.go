package help

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/wawandco/milo/cmd"
)

// Printer will print the help menu on the CLI.
type Printer struct {
	// Commands that the Printer will use to print the help.
	Commands []cmd.Runner
}

func (v Printer) Name() string {
	return "help"
}

// Run prints each of the Runners (in commands) registered help text if the command
// Is a cmd.HelpProvider.
func (v Printer) Run([]string) error {
	result := "Milo checks for issues with your HTML code.\n\n"
	result += "Usage:\n"
	result += "  milo [command] [args]\n\n"
	result += "Available Commands:"
	fmt.Print(result)

	// initialize tabwriter
	w := new(tabwriter.Writer)
	defer w.Flush()

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 3, '\t', 0)

	for _, command := range v.Commands {
		c, ok := command.(cmd.HelpProvider)
		if !ok {
			continue
		}

		fmt.Fprintf(w, "\n %v\t%v", c.Name(), c.HelpText())
	}

	// Printing help command
	fmt.Fprintf(w, "\n %v\t%v", v.Name(), v.HelpText())
	fmt.Fprintf(w, "\n")

	return nil
}

func (v Printer) HelpText() string {
	return "shows the help content for guidance."
}
