package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_InlineCSS_Review(t *testing.T) {
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
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, tcase.faultsLen, len(faults))
		if tcase.faultsLen == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.faults)

	}

}

func Test_InlineCSS_Accept(t *testing.T) {
	doc := reviewers.PageInlineCSSForbidden{}

	assert.True(t, doc.Accepts("/very/long/path/name/_partial.plush.html"), 
		"Expected to accept /very/long/path/name/_partial.plush.html")
	assert.True(t, doc.Accepts("_partial.plush.html"),
		"Expected to accept _partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected to accept templates/_partial.plush.html")
}
