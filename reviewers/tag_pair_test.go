package reviewers_test

import (
	"bytes"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_TagPair_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.TagPair{}
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
				Rule:     reviewers.Rules["0015"],
			}},
		},

		{
			name:      "single closed tag",
			faultsLen: 1,
			content:   "</html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			}},
		},

		{
			name:      "single open tag between paired tag",
			faultsLen: 1,
			content:   "<html><div></html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			}},
		},

		{
			name:      "single closed tag between paired tag",
			faultsLen: 1,
			content:   "<html></div></html>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			}},
		},

		{
			name:      "single open tag between paired tag followed by closed tag",
			faultsLen: 2,
			content:   "<html><div></html></div>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
				}},
		},

		{
			name:      "single closed tags followed by good matching tag",
			faultsLen: 3,
			content:   "<html></html></div></a></div>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
				}},
		},

		{
			name:      "single open tags followed by good matching tag",
			faultsLen: 3,
			content:   "<html></html><div><a><body>",
			fault: []reviewers.Fault{{
				Reviewer: reviewer.ReviewerName(),
				Line:     1,
				Rule:     reviewers.Rules["0015"],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
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
				Rule:     reviewers.Rules["0015"],
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
				Rule:     reviewers.Rules["0015"],
			},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Rule:     reviewers.Rules["0015"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Rule:     reviewers.Rules["0015"],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     6,
					Rule:     reviewers.Rules["0015"],
				}},
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
