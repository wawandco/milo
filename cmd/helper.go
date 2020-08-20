package cmd

type HelpProvider interface {
	Name() string
	HelpText() string
}
