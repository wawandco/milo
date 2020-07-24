package cmd

type Command interface {
	Run([]string) error
	CommandName() string
}
