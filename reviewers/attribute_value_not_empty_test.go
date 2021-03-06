package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrValueNotEmpty_Review(t *testing.T) {
	r := require.New(t)

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

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
		if tcase.faultsLen == 0 {
			continue
		}

		for i := range tcase.fault {
			r.Equal(faults[i].Reviewer, tcase.fault[i].Reviewer, tcase.name)
			r.Equal(faults[i].Line, tcase.fault[i].Line, tcase.name)
			r.Equal(faults[i].Col, tcase.fault[i].Col, tcase.name)
			r.Equal(faults[i].Rule.Code, tcase.fault[i].Rule.Code, tcase.name)
			r.Equal(faults[i].Rule.Description, tcase.fault[i].Rule.Description, tcase.name)
			r.Equal("something.html", faults[i].Path)
		}
	}

}
