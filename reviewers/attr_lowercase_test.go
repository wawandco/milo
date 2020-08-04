package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrLowercase_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrLowercase{}
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
				<img Src="test.png" />
				<input tYpe="image" />
				<area SHAPE="circle" COORDS="180,139,14" HREF="test.html" />
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Rule:     reviewers.Rules["0013"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Rule:     reviewers.Rules["0013"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Rule:     reviewers.Rules["0013"],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, len(tcase.faults), tcase.name)
		if len(tcase.faults) == 0 {
			continue
		}

		for i, tfault := range tcase.faults {
			r.Equal(faults[i].Reviewer, tfault.Reviewer, tcase.name)
			r.Equal(faults[i].Line, tfault.Line, tcase.name)
			r.Equal(faults[i].Rule.Code, tfault.Rule.Code, tcase.name)
			r.Equal(faults[i].Rule.Description, tfault.Rule.Description, tcase.name)
			r.Equal("something.html", faults[i].Path)
		}
	}

}
