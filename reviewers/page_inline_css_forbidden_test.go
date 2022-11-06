package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_InlineCSS_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.PageInlineCSSForbidden{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:      "no inline css",
			faultsLen: 0,
			content:   "<div></div>",
		},

		{
			name:      "inline div",
			faultsLen: 1,
			content:   `<div style="background-color: red;"></div>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      6,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "no inline div css tricky",
			faultsLen: 0,
			content:   `<div>style=something</div>`,
		},

		{
			name:      "uppercase attribute",
			faultsLen: 1,
			content:   `<div STYLE="something"></div>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      6,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "single quote",
			faultsLen: 1,
			content:   `<div STYLE='something'></div>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      6,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "self-closing",
			faultsLen: 1,
			content:   `<input style="some: css;" />`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      8,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "multiple inlines",
			faultsLen: 3,
			content: `
			<div style='something'>
				<div style='something'></div>
			</div>
			<hr style='something'>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      9,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      10,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      8,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoErr(err)
		r.Equal(len(faults), tcase.faultsLen)

		if tcase.faultsLen == 0 {
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

func Test_InlineCSS_Accept(t *testing.T) {
	r := is.New(t)
	doc := reviewers.PageInlineCSSForbidden{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}
