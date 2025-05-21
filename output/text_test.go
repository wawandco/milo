package output

import (
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_TextOutput(t *testing.T) {
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
	if got, want := formatter.FormatterName(), "text"; got != want {
		t.Errorf("expected formatter name %q, got %q", want, got)
	}

	out := formatter.Format(fault)
	expected := "file/does/not_exist.html:12:25: This is a test rule (1234:test/one)"
	if out != expected {
		t.Errorf("expected output %q, got %q", expected, out)
	}
}
