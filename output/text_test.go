package output

import (
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_TextOutput(t *testing.T) {
	r := is.New(t)

	fault := reviewers.Fault{
		Reviewer: "test/one",
		Line:     12,
		Col:      25,
		Path:     "file/does/not_exist.html",
		Rule: reviewers.Rule{
			Code:        "1234",
			Name:        "test-one",
			Description: "This is a test rule",
		},
	}

	formatter := TextFaultFormatter{}
	r.Equal(formatter.FormatterName(), "text")

	out := formatter.Format(fault)
	r.Equal("file/does/not_exist.html:12:25: This is a test rule (1234:test/one)", out)

}
