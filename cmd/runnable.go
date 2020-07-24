package cmd

type Runnable interface {
	Run([]string) error
	Name() string
}
