package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
	"github.com/wawandco/milo/reviewers"
)

func Test_TagPair_Review(t *testing.T) {
	

	reviewer := reviewers.PageTagParity{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		fault     []reviewers.Fault
	}{
		{
			name:      "empty",
			faultsLen: 0,
			content:   "",
		},

		{
			name:      "self-closed tag",
			faultsLen: 0,
			content:   "<br/>",
		},

		{
			name:      "tag paired",
			faultsLen: 0,
			content:   "<html></html>",
		},

		{
			name:      "single open tag",
			faultsLen: 1,
			content:   "<html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      1,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "single closed tag",
			faultsLen: 1,
			content:   "</html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      1,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "single open tag between paired tag",
			faultsLen: 1,
			content:   "<html><div></html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      7,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "single closed tag between paired tag",
			faultsLen: 1,
			content:   "<html></div></html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      7,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "single open tag between paired tag followed by closed tag",
			faultsLen: 2,
			content:   "<html><div></html></div>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      19,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      7,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},

		{
			name:      "single closed tags followed by good matching tag",
			faultsLen: 3,
			content:   "<html></html></div></a></div>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      14,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      20,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      24,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},

		{
			name:      "single open tags followed by good matching tag",
			faultsLen: 3,
			content:   "<html></html><div><a><body>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      14,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      19,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      22,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},

		{
			name:      "nested tags",
			faultsLen: 0,
			content:   "<ul><li><a></a></li></ul>",
		},

		{
			name:      "nested tags with invalid anchor closed",
			faultsLen: 1,
			content:   "<ul><li><a></a></a></li></ul>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Col:      16,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			}},
		},

		{
			name:      "more complex nested content with false positive",
			faultsLen: 4,
			content: `<ul>
						<li class="breadcrumb-item">
							<a href="#">Amenities</a>
						</li>
						<li class="breadcrumb-item active" aria-current="page">
							<span>Activity<span>
						</li>
						<li class="breadcrumb-item active" aria-current="page">
							<span>Edit Amenity</span>
						</li>
						<div>Bad child for ul</div>
						<span>Bad child for ul</span>
					 </ul>`,
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     13,
				Col:      7,
				Rule:     reviewers.Rules[reviewer.ReviewerName()],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      1,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Col:      8,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Col:      22,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				}},
		},
	}

	for _, tcase := range tcases {
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, tcase.faultsLen, len(faults))
		if tcase.faultsLen == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.fault)
	}

}
