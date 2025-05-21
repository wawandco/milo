package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_StyleAttribute_Review(t *testing.T) {
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
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, len(tcase.faults), len(faults))

		if len(tcase.faults) == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.faults)
	}

}

func Test_StyleAttribute_Accept(t *testing.T) {
	doc := reviewers.AttributeStyleForbidden{}

	assert.True(t, doc.Accepts("/very/long/path/name/_partial.plush.html"), 
		"Expected to accept /very/long/path/name/_partial.plush.html")
	assert.True(t, doc.Accepts("_partial.plush.html"),
		"Expected to accept _partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected to accept templates/_partial.plush.html")
}

func Test_StyleAttribute_Name(t *testing.T) {
	doc := reviewers.AttributeStyleForbidden{}
	assert.Equal(t, "attribute-style-forbidden", doc.ReviewerName())
}
