package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrIDUnique_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrIDUnique{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "no fault",
			content: `
				<img src="test.png" id="A"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
				<img src="test.png" id="D"/>
				<img src="test.png" id="b"/>
			`,
		},

		{
			name: "Faults",
			content: `
				<img src="test.png" id="A"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
				<img src="test.png" id="D"/>
				<img src="test.png" id="B"/>
				<img src="test.png" id="c"/>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     7,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
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
