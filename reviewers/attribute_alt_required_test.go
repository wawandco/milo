package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_AltRequired_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.AttributeAltRequired{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "no fault",
			content: `
				<img src="test.png" alt="test" />
				<input type="image" alt="test" />
				<area shape="circle" coords="180,139,14" href="test.html" alt="test" />
			`,
		},

		{
			name: "Faults",
			content: `
				<img src="test.png" />
				<input type="image" />
				<area shape="circle" coords="180,139,14" href="test.html" />
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name: "UPPERCASE",
			content: `
				<img src="test.png" ALT="test" />
				<input type="image" ALT="test" />
				<area shape="circle" coords="180,139,14" href="test.html" ALT="test" />
			`,
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoErr(err)
		r.Equal(len(faults), len(tcase.faults))
		if len(tcase.faults) == 0 {
			continue
		}

		for i, tfault := range tcase.faults {
			r.Equal(faults[i].Reviewer, tfault.Reviewer)
			r.Equal(faults[i].Line, tfault.Line)
			r.Equal(faults[i].Col, tfault.Col)
			r.Equal(faults[i].Rule.Code, tfault.Rule.Code)
			r.Equal(faults[i].Rule.Description, tfault.Rule.Description)
			r.Equal("something.html", faults[i].Path)
		}
	}

}
