package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_DoctypePresent_Review(t *testing.T) {
	doc := reviewers.PageDoctypePresent{}
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
					Col:      7,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name:    "no doctype",
			content: "<html></html>",
		},

		{
			name:    "partial should be omitted",
			content: `<div></div>`,
		},

		{
			faults: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     3,
					Col:      10,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name: "no doctype",
			content: `

			<html></html>
			`,
		},

		{
			faults: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Col:      17,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name: "no doctype",
			content: `<html lang="en"></html>
			`,
		},

		{
			faults: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Col:      17,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name: "uppercase",
			content: `<HTML lang="en"></HTML>
			`,
		},

		{
			name:    "sameline",
			content: `<!DOCTYPE html><html></html>`,
		},

		{
			name: "valid next line",
			content: `<!DOCTYPE html>
			<html>
			</html>`,
		},

		{
			name: "valid space line",
			content: `<!DOCTYPE html>

			<html>
			</html>`,
		},

		{
			name: "doctype case insensitive",
			content: `<!doctype html>
			<html lang="en">
			</html>`,
		},

		{
			name: "doctype old",
			content: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
			<html lang="en">
			</html>`,
		},

		{
			name: "no html tag",
			content: `
				<% contentFor("title") {%>
					Edit amenity
			  	<% } %>

				<%= contentFor("breadcrumb"){%>
					<nav aria-label="breadcrumb">
					<ol class="breadcrumb bg-none px-0 py-1">
						<li class="breadcrumb-item">
							<a href="<%= amenitiesPath() %>">Amenities</a>
						</li>
						<li class="breadcrumb-item active" aria-current="page">
							<span><%= amenity.Name %><span>
						</li>
						<li class="breadcrumb-item active" aria-current="page">
							<span>Edit Amenity<span>
						</li>
					</ol>
					</nav>
				<%} %>
			`,
		},

		{
			name: "xmlns case",
			content: `
				<?xml version="1.0"?>
				<!DOCTYPE html>
				<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
				<head>
					<title>Milo Test</title>
				</head>
				<body>
					<h1>Milo Test</h1>
				</body>
				</html>
			`,
		},

		{
			name: "php expression",
			content: `
				<!DOCTYPE html>
				<?php ?>
				<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
				<head>
					<title>Milo Test</title>
				</head>
				<body>
					<h1>Milo Test</h1>
				</body>
				</html>
			`,
		},

		{
			name: "comment",
			content: `
				<!DOCTYPE html>
				<!-- comment -->
				<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
				<head>
					<title>Milo Test</title>
				</head>
				<body>
					<h1>Milo Test</h1>
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
				t.Errorf("expected %v, got %v", faults[0].Path, "something.html")
			}
		}
	}
}

func Test_DoctypePresent_Accept(t *testing.T) {
	doc := reviewers.PageDoctypePresent{}

	if doc.Accepts("_partial.plush.html") {
		t.Error("Expected not to accept _partial.plush.html")
	}
	if !doc.Accepts("page.plush.html") {
		t.Error("Expected to accept page.plush.html")
	}
	if doc.Accepts("templates/_partial.plush.html") {
		t.Error("Expected not to accept templates/_partial.plush.html")
	}
}
