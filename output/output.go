package output

import "github.com/wawandco/milo/reviewers"

var Formatters = []FaultFormatter{
	GithubFaultFormatter{},
	SilentFaultFormatter{},
}

// FaultPrinter is intended to print faults in a specific format.
type FaultFormatter interface {
	FormatterName() string
	Format(reviewers.Fault) string
}
