package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
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

		page := bytes.NewBufferString("<a " + tcase.attr + " />")
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, tcase.faultsLen, len(faults))
		if tcase.faultsLen == 0 {
			continue
		}

		assert.Equal(t, reviewers.Rules[reviewer.ReviewerName()].Code, faults[0].Rule.Code)
		assert.Equal(t, reviewers.Rules[reviewer.ReviewerName()].Description, faults[0].Rule.Description)
		assert.Equal(t, reviewer.ReviewerName(), faults[0].Reviewer)
		assert.Equal(t, 4, faults[0].Col)
		assert.Equal(t, "something.html", faults[0].Path)
	}

}
