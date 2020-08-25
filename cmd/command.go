// package cmd holds common things for different executable commands in this package.
package cmd

// Runner interface is used for the commands that milo allows to run, these
// will match the following interface in order to be added to the possible commands.
type Runner interface {

	// Run method will receive the args, Run the command and
	// possibly return an error from the execution.
	Run([]string) error

	// Name returned here is used to identify the command that is
	// being invoked. p.e: milo review will call review.Runner.
	Name() string
}
