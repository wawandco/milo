package review

import flag "github.com/spf13/pflag"

var (
	// cliOutput holds the passed output flag value and determines the selected output format.
	cliOutput string
)

func init() {
	flag.StringVar(&cliOutput, "output", "", "the output format")
}
