package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_StyleAttribute_Review(t *testing.T) {
	reviewer := reviewers.AttributeStyleForbidden{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name:    "img full",
			content: `<img src="test.png" />`,
		},

		{
			name:    "img full",
			content: `<img src="test.png" style="something" /> <span style="something" />`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      21,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      48,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tcase.faults) != len(faults) {
			t.Errorf("expected length %d, got %d", len(tcase.faults), len(faults))
		}

		if len(tcase.faults) == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			if fault.Reviewer != faults[index].Reviewer {
				t.Errorf("expected %v, got %v", fault.Reviewer, faults[index].Reviewer)
			}
			if fault.Line != faults[index].Line {
				t.Errorf("expected %v, got %v", fault.Line, faults[index].Line)
			}
			if fault.Col != faults[index].Col {
				t.Errorf("expected %v, got %v", fault.Col, faults[index].Col)
			}
			if fault.Rule.Code != faults[index].Rule.Code {
				t.Errorf("expected %v, got %v", fault.Rule.Code, faults[index].Rule.Code)
			}
			if fault.Rule.Description != faults[index].Rule.Description {
				t.Errorf("expected %v, got %v", fault.Rule.Description, faults[index].Rule.Description)
			}
			if "something.html" != faults[index].Path {
				t.Errorf("expected %v, got %v", "something.html", faults[index].Path)
			}
		}

	}

}

func Test_StyleAttribute_Accept(t *testing.T) {
	doc := reviewers.AttributeStyleForbidden{}

	if !doc.Accepts("/very/long/path/name/_partial.plush.html") {
		t.Error("Expected to accept /very/long/path/name/_partial.plush.html")
	}
	if !doc.Accepts("_partial.plush.html") {
		t.Error("Expected to accept _partial.plush.html")
	}
	if !doc.Accepts("page.plush.html") {
		t.Error("Expected to accept page.plush.html")
	}
	if !doc.Accepts("templates/_partial.plush.html") {
		t.Error("Expected to accept templates/_partial.plush.html")
	}
}

func Test_StyleAttribute_Name(t *testing.T) {
	doc := reviewers.AttributeStyleForbidden{}
	if doc.ReviewerName() != "attribute-style-forbidden" {
		t.Errorf("expected %v, got %v", "attribute-style-forbidden", doc.ReviewerName())
	}
}
