package reviewers_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

func Test_SrcEmpty_Review(t *testing.T) {
	r := is.New(t)

	reviewer := reviewers.AttributeSrcRequired{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:      "img full",
			faultsLen: 0,
			content:   `<img src="test.png" />`,
		},

		{
			name:      "script full",
			faultsLen: 0,
			content:   `<script src="test.js"></script>`,
		},

		{
			name:      "link full",
			faultsLen: 0,
			content:   `<link href="test.css" type="text/css" />`,
		},

		{
			name:      "embed full",
			faultsLen: 0,
			content:   `<embed src="test.swf">`,
		},

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
					Col:      12,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "ignore comment",
			faultsLen: 0,
			content: `
			<html>
				<head></head>
				<body>
					<!-- <link href="" type="text/css"> -->
				</body>
			</html>
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

func Test_SrcEmpty_Accept(t *testing.T) {
	r := is.New(t)
	doc := reviewers.AttributeSrcRequired{}

	r.True(doc.Accepts("/very/long/path/name/_partial.plush.html"))
	r.True(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
	r.True(doc.Accepts("templates/_partial.plush.html"))
}

func Test_SrcEmpty_Name(t *testing.T) {
	r := is.New(t)
	doc := reviewers.AttributeSrcRequired{}
	r.Equal(doc.ReviewerName(), "attribute-src-required")
}
