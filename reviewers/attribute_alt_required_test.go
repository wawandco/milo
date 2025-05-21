package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_AltRequired_Review(t *testing.T) {

	reviewer := reviewers.AttributeAltRequired{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "no fault",
			content: `
				<img src="test.png" alt="test" />
				<input type="image" alt="test" />
				<area shape="circle" coords="180,139,14" href="test.html" alt="test" />
			`,
		},

		{
			name: "Faults",
			content: `
				<img src="test.png" />
				<input type="image" />
				<area shape="circle" coords="180,139,14" href="test.html" />
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name: "UPPERCASE",
			content: `
				<img src="test.png" ALT="test" />
				<input type="image" ALT="test" />
				<area shape="circle" coords="180,139,14" href="test.html" ALT="test" />
			`,
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(faults) != len(tcase.faults) {
			t.Fatalf("expected %d faults, got %d", len(tcase.faults), len(faults))
		}
		if len(tcase.faults) == 0 {
			continue
		}

		for i, tfault := range tcase.faults {
			if faults[i].Reviewer != tfault.Reviewer {
				t.Errorf("expected Reviewer %s, got %s", tfault.Reviewer, faults[i].Reviewer)
			}
			if faults[i].Line != tfault.Line {
				t.Errorf("expected Line %d, got %d", tfault.Line, faults[i].Line)
			}
			if faults[i].Col != tfault.Col {
				t.Errorf("expected Col %d, got %d", tfault.Col, faults[i].Col)
			}
			if faults[i].Rule.Code != tfault.Rule.Code {
				t.Errorf("expected Rule.Code %s, got %s", tfault.Rule.Code, faults[i].Rule.Code)
			}
			if faults[i].Rule.Description != tfault.Rule.Description {
				t.Errorf("expected Rule.Description %s, got %s", tfault.Rule.Description, faults[i].Rule.Description)
			}
			if faults[i].Path != "something.html" {
				t.Errorf("expected Path %s, got %s", "something.html", faults[i].Path)
			}
		}
	}

}
