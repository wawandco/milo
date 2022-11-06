package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_AttrIDUnique_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.AttributeIDUnique{}
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
			name: "no fault empty",
			content: `
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
				<img src="test.png" />
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
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     7,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
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
