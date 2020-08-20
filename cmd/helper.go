package cmd

type HelpProvider interface {
	Runner

	HelpText() string
}
