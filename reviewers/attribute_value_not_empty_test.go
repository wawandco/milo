package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
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

		assert.NoError(t, err)
		assert.Equal(t, tcase.faultsLen, len(faults))
		if tcase.faultsLen == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.fault)
	}

}
