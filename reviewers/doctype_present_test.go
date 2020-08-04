package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_DoctypePresent_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.DoctypePresent{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     reviewers.Fault
	}{
		{

			fault: reviewers.Fault{
				Reviewer: doc.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0001"],
			},
			name:      "no doctype",
			faultsLen: 1,
			content:   "<html></html>",
		},

		{
			name:      "partial should be omitted",
			faultsLen: 0,
			content:   `<div></div>`,
		},

		{
			fault: reviewers.Fault{
				Reviewer: doc.ReviewerName(),
				Line:     3,
				Rule:     reviewers.Rules["0001"],
			},
			name:      "no doctype",
			faultsLen: 1,
			content: `
			
			<html></html>
			`,
		},

		{
			fault: reviewers.Fault{
				Reviewer: doc.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0001"],
			},
			name:      "no doctype",
			faultsLen: 1,
			content: `<html lang="en"></html>
			`,
		},

		{
			fault: reviewers.Fault{
				Reviewer: doc.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0001"],
			},
			name:      "uppercase",
			faultsLen: 1,
			content: `<HTML lang="en"></HTML>
			`,
		},

		{
			name:      "sameline",
			faultsLen: 0,
			content:   `<!DOCTYPE html><html></html>`,
		},

		{
			name:      "valid next line",
			faultsLen: 0,
			content: `<!DOCTYPE html>
			<html>
			</html>`,
		},

		{
			name:      "valid space line",
			faultsLen: 0,
			content: `<!DOCTYPE html>

			<html>
			</html>`,
		},

		{
			name:      "doctype case insensitive",
			faultsLen: 0,
			content: `<!doctype html>
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
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := doc.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
		if tcase.faultsLen == 0 {
			continue
		}

		r.Equal(faults[0].Reviewer, tcase.fault.Reviewer, tcase.name)
		r.Equal(faults[0].Line, tcase.fault.Line, tcase.name)
		r.Equal(faults[0].Rule.Code, tcase.fault.Rule.Code, tcase.name)
		r.Equal(faults[0].Rule.Description, tcase.fault.Rule.Description, tcase.name)
		r.Equal("something.html", faults[0].Path)
	}

}

func Test_DoctypePresent_Accept(t *testing.T) {
	r := require.New(t)

	doc := reviewers.DoctypePresent{}

	r.False(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))

	r.False(doc.Accepts("templates/_partial.plush.html"))
}
