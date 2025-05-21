package reviewers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_AttrValueDoubleQuotes_Review(t *testing.T) {
	reviewer := reviewers.AttributeValueDoubleQuotes{}
	tcases := []struct {
		name      string
		attr      string
		faultsLen int
	}{
		{
			name:      "empty attribute should not be a fault",
			faultsLen: 0,
			attr:      "attr",
		},

		{
			name:      `attribute with " should not be a fault`,
			faultsLen: 0,
			attr:      `attr=""`,
		},

		{
			name:      "attribute with ' should be a fault",
			faultsLen: 1,
			attr:      "attr=''",
		},

		{
			name:      `attribute without " should be a fault`,
			faultsLen: 1,
			attr:      "attr=val",
		},

		{
			name:      `attribute without " should be a fault`,
			faultsLen: 1,
			attr:      "attr=",
		},
	}

	for _, tcase := range tcases {

		page := bytes.NewBufferString(fmt.Sprintf("<a %s />", tcase.attr))
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

		if faults[0].Rule.Code != reviewers.Rules[reviewer.ReviewerName()].Code {
			t.Errorf("expected %v, got %v", reviewers.Rules[reviewer.ReviewerName()].Code, faults[0].Rule.Code)
		}
		if faults[0].Rule.Description != reviewers.Rules[reviewer.ReviewerName()].Description {
			t.Errorf("expected %v, got %v", reviewers.Rules[reviewer.ReviewerName()].Description, faults[0].Rule.Description)
		}
		if faults[0].Reviewer != reviewer.ReviewerName() {
			t.Errorf("expected %v, got %v", reviewer.ReviewerName(), faults[0].Reviewer)
		}
		if faults[0].Col != 4 {
			t.Errorf("expected %v, got %v", 4, faults[0].Col)
		}
		if "something.html" != faults[0].Path {
			t.Errorf("expected %v, got %v", "something.html", faults[0].Path)
		}
	}

}
