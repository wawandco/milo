package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_TagLowercase_Review(t *testing.T) {
	r := require.New(t)

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
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, len(tcase.faults), tcase.name)

		if tcase.faultsLen == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(fault.Reviewer, faults[index].Reviewer, tcase.name)
			r.Equal(fault.Line, faults[index].Line, tcase.name)
			r.Equal(fault.Col, faults[index].Col, tcase.name)
			r.Equal(fault.Rule.Code, faults[index].Rule.Code, tcase.name)
			r.Equal(fault.Rule.Description, faults[index].Rule.Description, tcase.name)
			r.Equal("something.html", fault.Path)
		}

	}

}

func Test_TagLowercase_Accept(t *testing.T) {
	r := require.New(t)
	doc := reviewers.PageTagLowercaseRequired{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}

func Test_TagLowercase_Name(t *testing.T) {
	r := require.New(t)
	doc := reviewers.PageTagLowercaseRequired{}
	r.Equal(doc.ReviewerName(), "page-tag-lowercase")
}
