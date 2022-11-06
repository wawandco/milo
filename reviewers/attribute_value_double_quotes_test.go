package reviewers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_AttrValueDoubleQuotes_Review(t *testing.T) {
	r := is.New(t)

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

		r.NoErr(err)
		r.Equal(len(faults), tcase.faultsLen)
		if tcase.faultsLen == 0 {
			continue
		}

		r.Equal(faults[0].Rule.Code, reviewers.Rules[reviewer.ReviewerName()].Code)
		r.Equal(faults[0].Rule.Description, reviewers.Rules[reviewer.ReviewerName()].Description)
		r.Equal(faults[0].Reviewer, reviewer.ReviewerName())
		r.Equal(faults[0].Col, 4)
		r.Equal("something.html", faults[0].Path)
	}

}
