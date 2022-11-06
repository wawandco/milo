package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_StyleTag_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.PageStyleTagForbidden{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:      "no inline css",
			faultsLen: 0,
			content:   "<div></div>",
		},

		{
			name:      "style tag present in partial",
			faultsLen: 1,
			content:   `<div> <style class=""></style></div>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      7,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "style tag present full page",
			faultsLen: 1,
			content: `
			<html>
				<head><head>
				<body
					<div> <STYLE></STYLE></div>
				<body>
			<html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      12,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "style tag present in comment",
			faultsLen: 0,
			content: `
			<html>
				<head><head>
				<body
					<div> <!-- <STYLE></STYLE> --></div>
				<body>
			<html>
			`,
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoErr(err)
		r.Equal(len(faults), tcase.faultsLen)

		if tcase.faultsLen == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(fault.Reviewer, faults[index].Reviewer)
			r.Equal(fault.Line, faults[index].Line)
			r.Equal(fault.Col, faults[index].Col)
			r.Equal(fault.Rule.Code, faults[index].Rule.Code)
			r.Equal(fault.Rule.Description, faults[index].Rule.Description)
			r.Equal("something.html", faults[index].Path)
		}

	}

}

func Test_StyleTag_Accept(t *testing.T) {
	r := is.New(t)
	doc := reviewers.PageStyleTagForbidden{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}
