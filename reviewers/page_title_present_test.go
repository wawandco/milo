package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_TitlePresent_Review(t *testing.T) {
	doc := reviewers.PageTitlePresent{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{

			faults: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Col:      1,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name: "no title specified",
			content: `
			<html>
				<head></head>
			</html>`,
		},

		{

			faults: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Col:      1,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name: "empty title",
			content: `
			<html>
				<head><title></title></head>
			</html>`,
		},

		{
			name: "title specified",
			content: `
			<html>
				<head><title attr="something">Page Title</title></head>
			</html>`,
		},

		{
			name: "title specified uppercase",
			content: `
			<html>
				<head><TITLE attr="something">Page Title</TITLE></head>
			</html>`,
		},

		{
			name: "title tricky spaces specified uppercase",
			content: `
			<html>
				<head>


					<TITLE 
						attr="something">
						Page Title
					</TITLE>
				</head>
			</html>`,
		},

		{
			name: "partial without html/head",
			content: `
			<div>Some partial without html/head</div>
			`,
		},

		{
			name: "real case one",
			content: `
			<!DOCTYPE html>
			<html>
			
			<head>
			  <meta name="viewport" content="width=device-width, initial-scale=1">
			  <meta charset="utf-8">
			  <title>Housing Platform</title>
			  <%= stylesheetTag("application.css") %>
			  <meta name="csrf-param" content="authenticity_token" />
			  <meta name="csrf-token" content="<%= authenticity_token %>" />
			  
			  <%= partial("/partials/favicon.plush.html") %>
			</head>
			`,
		},

		{
			name: "",
			content: `
			<!DOCTYPE html>
			<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
			<head >
				<title>Page Demo</title>
			</head>
			<body>
			</body>
			</html>
			`,
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := doc.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, len(tcase.faults), len(faults))
		if len(tcase.faults) == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.faults)

	}

}

func Test_TitlePresent_Accept(t *testing.T) {
	doc := reviewers.PageTitlePresent{}

	assert.False(t, doc.Accepts("_partial.plush.html"),
		"Expected not to accept _partial.plush.html")
	assert.False(t, doc.Accepts("very/long/folder/length/_partial.plush.html"),
		"Expected not to accept very/long/folder/length/_partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("page.something.plush.html"),
		"Expected to accept page.something.plush.html")
	assert.True(t, doc.Accepts("page.html"),
		"Expected to accept page.html")
	assert.False(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected not to accept templates/_partial.plush.html")
}
