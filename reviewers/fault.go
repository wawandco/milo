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
		"::error file=%s,line=%d, col=0::[%s] %s (%s)",
		f.Path,
		f.Line,

		f.Rule.Code,
		f.Rule.Description,
		f.Reviewer,
	)
}
