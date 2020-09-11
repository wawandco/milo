package output

import "github.com/wawandco/milo/reviewers"

// SilentFaultFormatter does not print faults.
type SilentFaultFormatter struct{}

func (gp SilentFaultFormatter) FormatterName() string {
	return "silent"
}

func (gp SilentFaultFormatter) Format(f reviewers.Fault) string {
	return ""
}
