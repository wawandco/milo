package reviewers

import "fmt"

type Fault struct {
	Reviewer string
	Path     string
	Line     int

	Rule Rule
}

func (f Fault) String() string {
	return fmt.Sprintf(
		"%v:%v %v:%v (%v)",
		f.Path,
		f.Line,

		f.Rule.Code,
		f.Rule.Description,

		f.Reviewer,
	)
}
