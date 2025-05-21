package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_SrcEmpty_Review(t *testing.T) {

	reviewer := reviewers.AttributeSrcRequired{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:      "img full",
			faultsLen: 0,
			content:   `<img src="test.png" />`,
		},

		{
			name:      "script full",
			faultsLen: 0,
			content:   `<script src="test.js"></script>`,
		},

		{
			name:      "link full",
			faultsLen: 0,
			content:   `<link href="test.css" type="text/css" />`,
		},

		{
			name:      "embed full",
			faultsLen: 0,
			content:   `<embed src="test.swf">`,
		},

		{
			name:      "empty href",
			faultsLen: 1,
			content: `
			<html>
				<head></head>
				<body>
					<link href="" type="text/css">
				</body>
			</html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      12,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "ignore comment",
			faultsLen: 0,
			content: `
			<html>
				<head></head>
				<body>
					<!-- <link href="" type="text/css"> -->
				</body>
			</html>
			`,
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

func Test_SrcEmpty_Accept(t *testing.T) {
	doc := reviewers.AttributeSrcRequired{}

	assert.True(t, doc.Accepts("/very/long/path/name/_partial.plush.html"), 
		"Expected to accept /very/long/path/name/_partial.plush.html")
	assert.True(t, doc.Accepts("_partial.plush.html"),
		"Expected to accept _partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected to accept templates/_partial.plush.html")
}

func Test_SrcEmpty_Name(t *testing.T) {
	doc := reviewers.AttributeSrcRequired{}
	assert.Equal(t, "attribute-src-required", doc.ReviewerName())
}
