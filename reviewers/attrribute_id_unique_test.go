package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_AttrIDUnique_Review(t *testing.T) {
	

	reviewer := reviewers.AttributeIDUnique{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "no fault",
			content: `
				<img src="test.png" id="A"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
				<img src="test.png" id="D"/>
				<img src="test.png" id="b"/>
			`,
		},

		{
			name: "no fault empty",
			content: `
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
			`,
		},

		{
			name: "Faults",
			content: `
				<img src="test.png" id="A"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
				<img src="test.png" id="D"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     7,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
		if len(faults) != len(tcase.faults) {
		t.Errorf("expected length %d, got %d", len(tcase.faults), len(faults))
	}
		if len(tcase.faults) == 0 {
			continue
		}

		for i, tfault := range tcase.faults {
			if faults[i].Reviewer != tfault.Reviewer {
		t.Errorf("expected %v, got %v", tfault.Reviewer, faults[i].Reviewer)
	}
			if faults[i].Line != tfault.Line {
		t.Errorf("expected %v, got %v", tfault.Line, faults[i].Line)
	}
			if faults[i].Col != tfault.Col {
		t.Errorf("expected %v, got %v", tfault.Col, faults[i].Col)
	}
			if faults[i].Rule.Code != tfault.Rule.Code {
		t.Errorf("expected %v, got %v", tfault.Rule.Code, faults[i].Rule.Code)
	}
			if faults[i].Rule.Description != tfault.Rule.Description {
		t.Errorf("expected %v, got %v", tfault.Rule.Description, faults[i].Rule.Description)
	}
			if "something.html" != faults[i].Path {
		t.Errorf("expected %v, got %v", faults[i].Path, "something.html")
	}
		}
	}

}
