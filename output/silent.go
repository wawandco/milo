package output

import "wawandco/milo/reviewers"

type SilentFaultFormatter struct{}

func (gp SilentFaultFormatter) FormatterName() string {
	return "silent"
}

func (gp SilentFaultFormatter) Format(f reviewers.Fault) string {
	return ""
}