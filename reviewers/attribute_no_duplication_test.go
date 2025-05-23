package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_AttrNoDuplication_Review(t *testing.T) {
	

	reviewer := reviewers.AttributeNoDuplication{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:    "no attributes duplicated",
			content: `<img/><span>Text</span>`,
		},

		{
			name: "attributes duplicated on self closed tag",
			content: `<img src="/path-to-image.ext" alt="image" src="/path-to-image-alt.ext"/>
					  <span>Text</span>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      43,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name: "attributes duplicated on open and close tag",
			content: `<img/>
					  <span class="my-class" class="my-class-again">Text</span>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     2,
				Col:      31,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name: "attributes duplicated on both open/close and self-closed tags",
			content: `<img src="/path-to-image.ext" alt="image" src="/path-to-image-alt.ext"/>
					  <span class="my-class" class="my-class-again">Text</span>`,
			fault: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      43,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      31,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},

		{
			name:    "attributes duplicated on both open/close and self-closed tags",
			content: `<a href="/company/5eaf45f1-74ee-443b-9e17-e30949935cb0/users" class="list-group-item list-group-item-action ERB">`,
			fault:   []reviewers.Fault{},
		},

		{
			name:    "srcset and src",
			content: `<img class="logo" src="/assets/images/agnte_white_logo@2x" srcset="/assets/images/agnte_white_logo@2x.png 2x, /assets/images/agnte_white_logo.png 1x">`,
			fault:   []reviewers.Fault{},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, len(tcase.fault), len(faults))
		if len(tcase.fault) == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.fault)
	}

}
