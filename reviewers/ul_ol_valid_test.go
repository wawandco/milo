package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_OlValid_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.OlValid{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "no ol specified",
			faultsLen: 0,
			content: `
			<html>
				<body></body>
			</html>`,
		},

		{
			name:      "ol specified correctly",
			faultsLen: 0,
			content: `
			<ol>
				<li></li>
			</ol>
			`,
		},

		{
			fault: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0008"],
				},

				{
					Reviewer: doc.ReviewerName(),
					Line:     2,
					Rule:     reviewers.Rules["0008"],
				},
			},
			name:      "ol invalid",
			faultsLen: 2,
			content: `
			<ol>
				<label></label>
				<div></div>
			</ol>
			`,
		},

		{
			fault: []reviewers.Fault{
				{
					Reviewer: doc.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0008"],
				},
			},
			name:      "inner ol invalid",
			faultsLen: 1,
			content: `
			<ol>
				<li>
					<ol>
						<label></label>
					</ol>
				</li>
			</ol>
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

		for i := range tcase.fault {
			r.Equal(faults[i].Reviewer, tcase.fault[i].Reviewer, tcase.name)
			r.Equal(faults[i].Line, tcase.fault[i].Line, tcase.name)
			r.Equal(faults[i].Rule.Code, tcase.fault[i].Rule.Code, tcase.name)
			r.Equal(faults[i].Rule.Description, tcase.fault[i].Rule.Description, tcase.name)
		}
	}

}
