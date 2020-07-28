package output

import (
	"fmt"

	"github.com/wawandco/milo/reviewers"
)

type GithubFaultFormatter struct{}

func (gp GithubFaultFormatter) FormatterName() string {
	return "github"
}

func (gp GithubFaultFormatter) Format(f reviewers.Fault) string {
	return fmt.Sprintf(
		"::error file=%s,line=%d, col=0::[%s] %s (%s)",
		f.Path,
		f.Line,

		f.Rule.Code,
		f.Rule.Description,
		f.Reviewer,
	)
}
