package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_DoctypeValid(t *testing.T) {
	r := require.New(t)
	doc := reviewers.DoctypeValid{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     reviewers.Fault
	}{
		{
			fault: reviewers.Fault{
				Line:     1,
				Reviewer: doc.ReviewerName(),
				Rule:     reviewers.Rules["0002"],
			},
			name:      "doctype old",
			faultsLen: 1,
			content: `<!DOCTYPE INVALID>
			<html lang="en">
			</html>`,
		},

		{
			fault:     reviewers.Fault{},
			name:      "doctype valid",
			faultsLen: 0,
			content: `<!DOCTYPE html>
			<html lang="en">
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
		r.Equal(faults[0].Rule.Code, tcase.fault.Rule.Code, tcase.name)
		r.Equal(faults[0].Rule.Description, tcase.fault.Rule.Description)
		r.Equal(faults[0].Line, tcase.fault.Line, tcase.name)
	}

}
