package reviewers_test

import (
	"bytes"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrNoWhiteSpaces_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrNoWhiteSpaces{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "Correct",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{},
		},
		{
			name: "line 3 class attr has whitespace",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class ="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Rule:     reviewers.Rules["0019"],
				},
			},
		},
		{
			name: "line 3 class attr has whitespace and line 4 data-block has space as well",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class ="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light" data-block= "last-name">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Rule:     reviewers.Rules["0019"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Rule:     reviewers.Rules["0019"],
				},
			},
		},
		{
			name: "whitespaces between uppercase attr and value",
			content: `
			<span CLASS=  "font-18 font-weight-bold">First Name</span>
			<p class="font-14 text-yellow" Name    ="LastName">LastName</p>
			<span class="font-12 text-muted font-weight-light" DATA-BLOCK  = "last-name">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Rule:     reviewers.Rules["0019"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Rule:     reviewers.Rules["0019"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Rule:     reviewers.Rules["0019"],
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
		}
	}
}
