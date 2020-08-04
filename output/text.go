package output

import (
	"fmt"

	"github.com/wawandco/milo/reviewers"
)

type TextFaultFormatter struct{}

func (gp TextFaultFormatter) FormatterName() string {
	return "text"
}

func (gp TextFaultFormatter) Format(f reviewers.Fault) string {
	return fmt.Sprintf(
		"%v:%v:1: %v (%v:%v)",
		f.Path,
		f.Line,
		f.Rule.Description,
		f.Rule.Code,
		f.Reviewer,
	)
}
