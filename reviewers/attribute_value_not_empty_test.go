package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_AttrValueNotEmpty_Review(t *testing.T) {
	

	reviewer := reviewers.AttributeValueNotEmpty{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "no attributes specified",
			faultsLen: 0,
			content:   `<img/><span>Text</span>`,
		},

		{
			name:      "attributes with values",
			faultsLen: 0,
			content:   `<img src="/path-to-image.ext" alt="image"/>`,
		},

		{
			name:      "checked attributte is valid",
			faultsLen: 0,
			content:   `<input type="ckeckbox" checked/>`,
			fault:     []reviewers.Fault{},
		},

		{
			name:      "disabled attributte is valid",
			faultsLen: 0,
			content:   `<input type="ckeckbox" disabled/>`,
			fault:     []reviewers.Fault{},
		},

		{
			name:      "one attribute with empty value",
			faultsLen: 1,
			content:   `<img src="" alt="image"/>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      6,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "all attributes with empty values",
			faultsLen: 2,
			content:   `<img src="" alt=""/>`,
			fault: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      6,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      13,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
		if len(faults) != tcase.faultsLen {
		t.Errorf("expected %v, got %v", tcase.faultsLen, len(faults))
	}
		if tcase.faultsLen == 0 {
			continue
		}

		for i := range tcase.fault {
			if faults[i].Reviewer != tcase.fault[i].Reviewer {
		t.Errorf("expected %v, got %v", tcase.fault[i].Reviewer, faults[i].Reviewer)
	}
			if faults[i].Line != tcase.fault[i].Line {
		t.Errorf("expected %v, got %v", tcase.fault[i].Line, faults[i].Line)
	}
			if faults[i].Col != tcase.fault[i].Col {
		t.Errorf("expected %v, got %v", tcase.fault[i].Col, faults[i].Col)
	}
			if faults[i].Rule.Code != tcase.fault[i].Rule.Code {
		t.Errorf("expected %v, got %v", tcase.fault[i].Rule.Code, faults[i].Rule.Code)
	}
			if faults[i].Rule.Description != tcase.fault[i].Rule.Description {
		t.Errorf("expected %v, got %v", tcase.fault[i].Rule.Description, faults[i].Rule.Description)
	}
			if "something.html" != faults[i].Path {
		t.Errorf("expected %v, got %v", faults[i].Path, "something.html")
	}
		}
	}

}
