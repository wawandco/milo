package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_StyleTag_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.StyleTag{}
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
					Rule:     reviewers.Rules["0005"],
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
					Rule:     reviewers.Rules["0005"],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)

		if tcase.faultsLen == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			r.Equal(fault.Reviewer, faults[index].Reviewer, tcase.name)
			r.Equal(fault.Line, faults[index].Line, tcase.name)
			r.Equal(fault.Rule.Code, faults[index].Rule.Code, tcase.name)
			r.Equal(fault.Rule.Description, faults[index].Rule.Description, tcase.name)
		}

	}

}

func Test_StyleTag_Accept(t *testing.T) {
	r := require.New(t)
	doc := reviewers.StyleTag{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}
