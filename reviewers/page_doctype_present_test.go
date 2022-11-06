package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_DoctypePresent_Review(t *testing.T) {
	r := is.New(t)

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

		r.NoErr(err)
		r.Equal(len(faults), len(tcase.faults))

		for index, fault := range tcase.faults {
			r.Equal(faults[index].Reviewer, fault.Reviewer)
			r.Equal(faults[index].Line, fault.Line)
			r.Equal(faults[index].Col, fault.Col)
			r.Equal(faults[index].Rule.Code, fault.Rule.Code)
			r.Equal(faults[index].Rule.Description, fault.Rule.Description)
			r.Equal("something.html", faults[0].Path)
		}
	}

}

func Test_DoctypePresent_Accept(t *testing.T) {
	r := is.New(t)

	doc := reviewers.PageDoctypePresent{}

	r.True(!doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))

	r.True(!doc.Accepts("templates/_partial.plush.html"))
}
