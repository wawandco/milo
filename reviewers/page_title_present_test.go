package reviewers_test

import (
	"strings"
	"testing"

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
		page := strings.NewReader(tcase.content)
		faults, err := doc.Review("something.html", page)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(faults) != len(tcase.faults) {
			t.Errorf("expected length %d, got %d", len(tcase.faults), len(faults))
		}
		if len(tcase.faults) == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			if faults[index].Reviewer != fault.Reviewer {
				t.Errorf("expected %v, got %v", fault.Reviewer, faults[index].Reviewer)
			}
			if faults[index].Line != fault.Line {
				t.Errorf("expected %v, got %v", fault.Line, faults[index].Line)
			}
			if faults[index].Col != fault.Col {
				t.Errorf("expected %v, got %v", fault.Col, faults[index].Col)
			}
			if faults[index].Rule.Code != fault.Rule.Code {
				t.Errorf("expected %v, got %v", fault.Rule.Code, faults[index].Rule.Code)
			}
			if faults[index].Rule.Description != fault.Rule.Description {
				t.Errorf("expected %v, got %v", fault.Rule.Description, faults[index].Rule.Description)
			}
			if "something.html" != faults[0].Path {
				t.Errorf("expected %v, got %v", "something.html", faults[0].Path)
			}
		}

	}

}

func Test_TitlePresent_Accept(t *testing.T) {
	doc := reviewers.PageTitlePresent{}

	if doc.Accepts("_partial.plush.html") {
		t.Error("Expected not to accept _partial.plush.html")
	}
	if doc.Accepts("very/long/folder/length/_partial.plush.html") {
		t.Error("Expected not to accept very/long/folder/length/_partial.plush.html")
	}
	if !doc.Accepts("page.plush.html") {
		t.Error("Expected to accept page.plush.html")
	}
	if !doc.Accepts("page.something.plush.html") {
		t.Error("Expected to accept page.something.plush.html")
	}
	if !doc.Accepts("page.html") {
		t.Error("Expected to accept page.html")
	}
	if doc.Accepts("templates/_partial.plush.html") {
		t.Error("Expected not to accept templates/_partial.plush.html")
	}
}
