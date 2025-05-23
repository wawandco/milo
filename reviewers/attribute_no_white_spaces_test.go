package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_AttrNoWhiteSpaces_Review(t *testing.T) {
	

	reviewer := reviewers.AttributeNoWhiteSpaces{}
	tcases := []struct {
		name    string
		content string
		err     error
		faults  []reviewers.Fault
	}{
		{
			name: "Correct",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{},
		},
		{
			name: "line 3 class attr has whitespace",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class ="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
		{
			name: "line 3 data attr with number has whitespace",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p data-1   = "value">LastName</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
		{
			name: "line 3 class attr has whitespace and line 4 data-block has space as well",
			content: `
				<span class="font-18 font-weight-bold">First Name</span>
				<p class ="font-14 text-yellow">LastName</p>
				<span class="font-12 text-muted font-weight-light" data-block= "last-name">Status</span>
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
					Line:     4,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
		{
			name: "whitespaces between uppercase attr and value",
			content: `
			<span CLASS=  "font-18 font-weight-bold">First Name</span>
			<p class="font-14 text-yellow" Name    ="LastName">LastName</p>
			<span class="font-12 text-muted font-weight-light" DATA-BLOCK  = "last-name">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      4,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      4,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Col:      4,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
		{
			name: "attribute with number",
			content: `
				<span class="font-18 font-weight-bold" data-2-Line= "26 north">custom Value</span>
				<p class="font-14 text-yellow">second value</p>
				<span class="font-12 text-muted font-weight-light">Status</span>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
		{
			name: "cover attributes with special characters",
			content: `
				<div data-field_status ="Valid">Valid Status</div>
				<div aria-label**name ="true">Label</div>
				<div data-attr/value =   "true">value</div>
			`,

			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     2,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     4,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, len(tcase.faults), len(faults))
		if len(tcase.faults) == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.faults)
	}
}
