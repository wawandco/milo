// Output package contains different output formats for the
// result of running different results.
package output

import "github.com/wawandco/milo/reviewers"

// Formatters available for use by the CLI.
var Formatters = []FaultFormatter{
	GithubFaultFormatter{},
	SilentFaultFormatter{},
	TextFaultFormatter{},
}

// FaultPrinter is intended to print faults in a specific format.
type FaultFormatter interface {
	FormatterName() string
	Format(reviewers.Fault) string
}
