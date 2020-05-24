package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_TitlePresent_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.TitlePresent{}
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
				Rule:     reviewers.Rules["0004"],
			},
			name:      "no title specified",
			faultsLen: 1,
			content: `
			<html>
				<head></head>
			</html>`,
		},

		{

			fault: reviewers.Fault{
				Reviewer: doc.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0004"],
			},
			name:      "empty title",
			faultsLen: 1,
			content: `
			<html>
				<head><title></title></head>
			</html>`,
		},

		{
			name:      "title specified",
			faultsLen: 0,
			content: `
			<html>
				<head><title attr="something">Page Title</title></head>
			</html>`,
		},

		{
			name:      "title specified uppercase",
			faultsLen: 0,
			content: `
			<html>
				<head><TITLE attr="something">Page Title</TITLE></head>
			</html>`,
		},

		{
			name:      "title tricky spaces specified uppercase",
			faultsLen: 0,
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
	}

}

func Test_TitlePresent_Accept(t *testing.T) {
	r := require.New(t)

	doc := reviewers.TitlePresent{}

	r.False(doc.Accepts("_partial.plush.html"))
	r.False(doc.Accepts("very/long/folder/length/_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("page.something.plush.html"))
	r.True(doc.Accepts("page.html"))
	r.False(doc.Accepts("templates/_partial.plush.html"))
}
