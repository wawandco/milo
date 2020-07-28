package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrNoDuplication_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrNoDuplication{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "no attributes duplicated",
			faultsLen: 0,
			content:   `<img/><span>Text</span>`,
		},

		{
			name:      "attributes duplicated on self closed tag",
			faultsLen: 1,
			content: `<img src="/path-to-image.ext" alt="image" src="/path-to-image-alt.ext"/>
					  <span>Text</span>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0010"],
			}},
		},

		{
			name:      "attributes duplicated on open and close tag",
			faultsLen: 1,
			content: `<img/>
					  <span class="my-class" class="my-class-again">Text</span>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     2,
				Rule:     reviewers.Rules["0010"],
			}},
		},

		{
			name:      "attributes duplicated on both open/close and self-closed tags",
			faultsLen: 2,
			content: `<img src="/path-to-image.ext" alt="image" src="/path-to-image-alt.ext"/>
					  <span class="my-class" class="my-class-again">Text</span>`,
			fault: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0010"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Rule:     reviewers.Rules["0010"],
				}},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
		if tcase.faultsLen == 0 {
			continue
		}

		for i := range tcase.fault {
			r.Equal(faults[i].Reviewer, tcase.fault[i].Reviewer, tcase.name)
			r.Equal(faults[i].Line, tcase.fault[i].Line, tcase.name)
			r.Equal(faults[i].Rule.Code, tcase.fault[i].Rule.Code, tcase.name)
			r.Equal(faults[i].Rule.Description, tcase.fault[i].Rule.Description, tcase.name)
		}
	}

}
