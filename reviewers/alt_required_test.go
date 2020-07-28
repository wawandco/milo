package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AltRequired_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AltRequired{}
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
					Rule:     reviewers.Rules["0012"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Rule:     reviewers.Rules["0012"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Rule:     reviewers.Rules["0012"],
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
		}
	}

}
