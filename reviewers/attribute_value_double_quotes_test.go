package reviewers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrValueDoubleQuotes_Review(t *testing.T) {
	r := require.New(t)

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

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
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
