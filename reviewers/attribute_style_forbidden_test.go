package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_StyleAttribute_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.AttributeStyleForbidden{}
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
					Col:      21,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      48,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoErr(err)
		r.Equal(len(tcase.faults), len(faults))

		if len(tcase.faults) == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(fault.Reviewer, faults[index].Reviewer)
			r.Equal(fault.Line, faults[index].Line)
			r.Equal(fault.Col, faults[index].Col)
			r.Equal(fault.Rule.Code, faults[index].Rule.Code)
			r.Equal(fault.Rule.Description, faults[index].Rule.Description)
			r.Equal("something.html", faults[index].Path)
		}

	}

}

func Test_StyleAttribute_Accept(t *testing.T) {
	r := is.New(t)
	doc := reviewers.AttributeStyleForbidden{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}

func Test_StyleAttribute_Name(t *testing.T) {
	r := is.New(t)
	doc := reviewers.AttributeStyleForbidden{}
	r.Equal(doc.ReviewerName(), "attribute-style-forbidden")
}
