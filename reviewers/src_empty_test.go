package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_SrcEmpty_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.SrcEmpty{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		// {
		// 	name:      "img full",
		// 	faultsLen: 0,
		// 	content:   `<img src="test.png" />`,
		// },

		// {
		// 	name:      "script full",
		// 	faultsLen: 0,
		// 	content:   `<script src="test.js"></script>`,
		// },

		// {
		// 	name:      "link full",
		// 	faultsLen: 0,
		// 	content:   `<link href="test.css" type="text/css" />`,
		// },

		// {
		// 	name:      "embed full",
		// 	faultsLen: 0,
		// 	content:   `<embed src="test.swf">`,
		// },

		{
			name:      "empty href",
			faultsLen: 1,
			content: `
			<html>
				<head></head>
				<body>
					<link href="" type="text/css">
				</body>
			</html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Rule:     reviewers.Rules["0007"],
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

func Test_SrcEmpty_Accept(t *testing.T) {
	r := require.New(t)
	doc := reviewers.SrcEmpty{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}

func Test_SrcEmpty_Name(t *testing.T) {
	r := require.New(t)
	doc := reviewers.SrcEmpty{}
	r.Equal(doc.ReviewerName(), "tag/src-empty")
}
