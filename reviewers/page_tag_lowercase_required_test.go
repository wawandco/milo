package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_TagLowercase_Review(t *testing.T) {
	reviewer := reviewers.PageTagLowercaseRequired{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:    "no inline css",
			content: "<div></div>",
		},

		{
			name:    "lowercased",
			content: `<div><style class=""></style></div>`,
		},

		{
			name: "UPPER case tag",
			content: `
			<html>
				<head><head>
				<body>
					<div><STYLE></STYLE></div>
				<body>
			<html>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      18,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name: "mixed cases on tag name",
			content: `
			<html>
				<Head><Head>
				<body
					<div><STYLE></STYLE></div>
				<body>
			<html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      18,
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

func Test_TagLowercase_Accept(t *testing.T) {
	doc := reviewers.PageTagLowercaseRequired{}

	assert.True(t, doc.Accepts("/very/long/path/name/_partial.plush.html"), 
		"Expected to accept /very/long/path/name/_partial.plush.html")
	assert.True(t, doc.Accepts("_partial.plush.html"),
		"Expected to accept _partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected to accept templates/_partial.plush.html")
}

func Test_TagLowercase_Name(t *testing.T) {
	doc := reviewers.PageTagLowercaseRequired{}
	assert.Equal(t, "page-tag-lowercase", doc.ReviewerName())
}
