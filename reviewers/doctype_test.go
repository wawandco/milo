package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_DoctypeReviewer_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.Doctype{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     reviewers.Fault
	}{
		{

			fault: reviewers.Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   1,
			},
			name:      "no doctype",
			faultsLen: 1,
			content:   "<html></html>",
		},

		{
			fault:     reviewers.Fault{ReviewerName: doc.ReviewerName()},
			name:      "partial should be omitted",
			faultsLen: 0,
			content:   `<div></div>`,
		},

		{
			fault: reviewers.Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   3,
			},
			name:      "no doctype",
			faultsLen: 1,
			content: `
			
			<html></html>
			`,
		},

		{
			fault: reviewers.Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   1,
			},
			name:      "no doctype",
			faultsLen: 1,
			content: `<html lang="en"></html>
			`,
		},

		{
			fault: reviewers.Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   1,
			},
			name:      "uppercase",
			faultsLen: 1,
			content: `<HTML lang="en"></HTML>
			`,
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := doc.Review(page)

		r.NoError(err, tcase.name)
		r.Len(faults, tcase.faultsLen, tcase.name)
		if tcase.faultsLen == 0 {
			continue
		}

		r.Equal(faults[0].ReviewerName, tcase.fault.ReviewerName, tcase.name)
		r.Equal(faults[0].LineNumber, tcase.fault.LineNumber, tcase.name)
	}

}

func Test_DoctypeReviewer_Accept(t *testing.T) {
	r := require.New(t)

	doc := reviewers.Doctype{}

	r.False(doc.Accepts("_partial.plush.html"))
	r.True(doc.Accepts("page.plush.html"))
}
