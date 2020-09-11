package review

import flag "github.com/spf13/pflag"

var (
	//
	cliOutput string
)

func init() {
	flag.StringVar(&cliOutput, "output", "", "the output format")
}
