package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_StyleAttribute_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.StyleAttribute{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name:    "img full",
			content: `<img src="test.png" />`,
		},

		{
			name:    "img full",
			content: `<img src="test.png" style="something" /> <span style="something" />`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0009"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0009"],
				},
			},
		},
	}

	for index, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(tcase.faults, len(faults), tcase.name, "Case %v", index+1)

		if len(tcase.faults) == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(fault.Reviewer, faults[index].Reviewer, tcase.name)
			r.Equal(fault.Line, faults[index].Line, tcase.name)
			r.Equal(fault.Rule.Code, faults[index].Rule.Code, tcase.name)
			r.Equal(fault.Rule.Description, faults[index].Rule.Description, tcase.name)
		}

	}

}

func Test_StyleAttribute_Accept(t *testing.T) {
	r := require.New(t)
	doc := reviewers.StyleAttribute{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}

func Test_StyleAttribute_Name(t *testing.T) {
	r := require.New(t)
	doc := reviewers.StyleAttribute{}
	r.Equal(doc.ReviewerName(), "attribute/style")
}
