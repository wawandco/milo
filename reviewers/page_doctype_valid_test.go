package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_DoctypeValid(t *testing.T) {
	
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

		if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
		if len(faults) != tcase.faultsLen {
		t.Errorf("expected %v, got %v", tcase.faultsLen, len(faults))
	}
		if tcase.faultsLen == 0 {
			continue
		}

		if faults[0].Reviewer != tcase.fault.Reviewer {
		t.Errorf("expected %v, got %v", tcase.fault.Reviewer, faults[0].Reviewer)
	}
		if faults[0].Rule.Code != tcase.fault.Rule.Code {
		t.Errorf("expected %v, got %v", tcase.fault.Rule.Code, faults[0].Rule.Code)
	}
		if faults[0].Rule.Description != tcase.fault.Rule.Description {
		t.Errorf("expected %v, got %v", tcase.fault.Rule.Description, faults[0].Rule.Description)
	}
		if faults[0].Line != tcase.fault.Line {
		t.Errorf("expected %v, got %v", tcase.fault.Line, faults[0].Line)
	}
		if faults[0].Col != tcase.fault.Col {
		t.Errorf("expected %v, got %v", tcase.fault.Col, faults[0].Col)
	}
		if "something.html" != faults[0].Path {
		t.Errorf("expected %v, got %v", faults[0].Path, "something.html")
	}
	}

}
