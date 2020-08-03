package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrValueNotEmpty_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrValueNotEmpty{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "no attributes specified",
			faultsLen: 0,
			content:   `<img/><span>Text</span>`,
		},

		{
			name:      "attributes with values",
			faultsLen: 0,
			content:   `<img src="/path-to-image.ext" alt="image"/>`,
		},

		{
			name:      "attribute with erb expression into attr value",
			faultsLen: 0,
			content:   `<a href="/company/5eaf45f1-74ee-443b-9e17-e30949935cb0/users" class="list-group-item list-group-item-action <%= activePathClass("active", "/company/.*/users/.*") %>">`,
		},

		{
			name:      "attribute with erb expression condition to show an attr value",
			faultsLen: 0,
			content:   `<a href="/company/5eaf45f1-74ee-443b-9e17-e30949935cb0/users" class="list-group-item list-group-item-action <%= if (ctVariable.Value != "CUSTOM_VALUE" && secondContition == 123456) { %>custom-class<% } %>">`,
		},

		{
			name:      "one attribute with empty value",
			faultsLen: 1,
			content:   `<img src="" alt="image"/>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0011"],
			}},
		},

		{
			name:      "all attributes with empty values",
			faultsLen: 2,
			content:   `<img src="" alt=""/>`,
			fault: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0011"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0011"],
				}},
		},

		{
			name:      "attribute with erb expression condition to show an attr",
			faultsLen: 1,
			content:   `<a href="/company/5eaf45f1-74ee-443b-9e17-e30949935cb0/users" class="list-group-item list-group-item-action" <%= if (ctVariable.Value != "CUSTOM_VALUE" && secondContition == 123456) { %>data-attr="" <% } %>>`,
			fault: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0011"],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

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
