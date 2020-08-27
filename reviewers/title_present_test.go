package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_TitlePresent_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.TitlePresent{}
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
		page := strings.NewReader(tcase.content)
		faults, err := doc.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, len(tcase.faults), tcase.name)
		if len(tcase.faults) == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(faults[index].Reviewer, fault.Reviewer, tcase.name)
			r.Equal(faults[index].Line, fault.Line, tcase.name)
			r.Equal(faults[index].Col, fault.Col, tcase.name)
			r.Equal(faults[index].Rule.Code, fault.Rule.Code, tcase.name)
			r.Equal(faults[index].Rule.Description, fault.Rule.Description, tcase.name)
			r.Equal("something.html", faults[0].Path)
		}

	}

}

func Test_TitlePresent_Accept(t *testing.T) {
	r := require.New(t)

	doc := reviewers.TitlePresent{}

	r.False(doc.Accepts("_partial.plush.html"))
	r.False(doc.Accepts("very/long/folder/length/_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("page.something.plush.html"))
	r.True(doc.Accepts("page.html"))
	r.False(doc.Accepts("templates/_partial.plush.html"))
}
