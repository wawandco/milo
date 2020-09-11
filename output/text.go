package output

import (
	"fmt"

	"github.com/wawandco/milo/reviewers"
)

// TextFaultFormatter prints faults in a simple way
// this way is useful for development environments because
// it allows to click and jump to the file.
type TextFaultFormatter struct{}

func (gp TextFaultFormatter) FormatterName() string {
	return "text"
}

func (gp TextFaultFormatter) Format(f reviewers.Fault) string {
	return fmt.Sprintf(
		"%v:%v:%v: %v (%v:%v)",
		f.Path,
		f.Line,
		f.Col,
		f.Rule.Description,
		f.Rule.Code,
		f.Reviewer,
	)
}
