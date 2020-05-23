package reviewers

import "fmt"

type Fault struct {
	ReviewerName string
	Path         string
	LineNumber   int

	Rule Rule
}

func (f Fault) String() string {
	return fmt.Sprintf(
		"%v:%v %v:%v (%v)",
		f.Path,
		f.LineNumber,
		f.Rule.Code,
		f.Rule.Description,
		f.ReviewerName,
	)
}
