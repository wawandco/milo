package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_DoctypeValid(t *testing.T) {
	r := is.New(t)
	doc := reviewers.PageDoctypeValid{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     reviewers.Fault
	}{
		{
			fault: reviewers.Fault{
				Line:     1,
				Col:      10,
				Reviewer: doc.ReviewerName(),
				Rule:     reviewers.Rules[doc.ReviewerName()],
			},
			name:      "doctype old",
			faultsLen: 1,
			content: `<!DOCTYPE INVALID>
			<html lang="en">
			</html>`,
		},

		{
			name:      "doctype old",
			faultsLen: 0,
			content: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
			<html lang="en">
			</html>`,
		},

		{
			fault:     reviewers.Fault{},
			name:      "doctype valid",
			faultsLen: 0,
			content: `<!DOCTYPE html>
			<html lang="en">
			</html>`,
		},

		{
			name:      "no html tag",
			faultsLen: 0,
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
			name:      "valid doctype real case",
			faultsLen: 0,
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
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := doc.Review("something.html", page)

		r.NoErr(err)
		r.Equal(len(faults), tcase.faultsLen)
		if tcase.faultsLen == 0 {
			continue
		}

		r.Equal(faults[0].Reviewer, tcase.fault.Reviewer)
		r.Equal(faults[0].Rule.Code, tcase.fault.Rule.Code)
		r.Equal(faults[0].Rule.Description, tcase.fault.Rule.Description)
		r.Equal(faults[0].Line, tcase.fault.Line)
		r.Equal(faults[0].Col, tcase.fault.Col)
		r.Equal("something.html", faults[0].Path)
	}

}
