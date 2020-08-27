package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_OlUlValid_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.OlUlValid{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "no ol/ul specified",
			faultsLen: 0,
			content: `
			<html>
				<body></body>
			</html>`,
		},

		{
			name:      "ol/ul specified correctly",
			faultsLen: 0,
			content: `
			<ol>
				<li></li>
			</ol>
			<ul>
				<li></li>
			</ul>
			`,
		},

		{
			fault: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},

				{
					Reviewer: doc.ReviewerName(),
					Line:     4,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
				{
					Reviewer: doc.ReviewerName(),
					Line:     7,
					Col:      5,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},

				{
					Reviewer: doc.ReviewerName(),
					Line:     8,
					Col:      5,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},
			name:      "ol/ul invalid",
			faultsLen: 4,
			content: `
			<ol>
				<label></label>
				<div></div>
			</ol>
			<ul>
				<label></label>
				<div></div>
			</ul>
			`,
		},

		{
			fault: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     5,
					Col:      7,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
				{
					Reviewer: doc.ReviewerName(),
					Line:     13,
					Col:      7,
					Rule:     reviewers.Rules[doc.ReviewerName()],
				},
			},

			name:      "inner ol/ul invalid",
			faultsLen: 2,
			content: `
			<ol>
				<li>
					<ol>
						<label></label>
					</ol>
				</li>
			</ol>

			<ul>
				<li>
					<ul>
						<label></label>
					</ul>
				</li>
			</ul>
			`,
		},

		{
			name:      "reported case",
			faultsLen: 0,
			content: `
				<!DOCTYPE html>
				<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
				<head>
					<title>Contact Us</title>
				</head>
				<body>
					<ul class="mainMenu nav nav-pills">
						<li><i/> Home</li>
					</ul>
				</body>
				</html>
			`,
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := doc.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
		if tcase.faultsLen == 0 {
			continue
		}

		for i := range tcase.fault {
			r.Equal(faults[i].Reviewer, tcase.fault[i].Reviewer, tcase.name)
			r.Equal(faults[i].Line, tcase.fault[i].Line, tcase.name)
			r.Equal(faults[i].Rule.Code, tcase.fault[i].Rule.Code, tcase.name)
			r.Equal(faults[i].Rule.Description, tcase.fault[i].Rule.Description, tcase.name)
			r.Equal("something.html", faults[i].Path)
		}
	}

}
