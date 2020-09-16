package cmd

// HelpProvider interface is implemented by commands that provide help text.
type HelpProvider interface {
	// Name is used when displaying help text in the table, it allows to associate help texts with the command.
	Name() string

	// HelpText returns the help text of the subcommand invoked.
	HelpText() string
}
