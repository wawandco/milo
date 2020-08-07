package reviewers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_AttrUnsafeChars_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.AttrUnsafeChars{}
	tcases := []struct {
		name      string
		chars     []string
		faultsLen int
	}{
		{
			name:      "no unsafe chars",
			faultsLen: 0,
			chars:     []string{"#"},
		},

		{
			name:      "null",
			faultsLen: 1,
			chars:     []string{"\u0000"},
		},

		{
			name:      "tabulation",
			faultsLen: 1,
			chars:     []string{"\u0009"},
		},

		{
			name:      "between null and tabulation",
			faultsLen: 1,
			chars: []string{"\u0001", "\u0002", "\u0003", "\u0004",
				"\u0005", "\u0006", "\u0007", "\u0008"},
		},

		{
			name:      "line Tabulation",
			faultsLen: 1,
			chars:     []string{"\u000b"},
		},

		{
			name:      "form feed",
			faultsLen: 1,
			chars:     []string{"\u000c"},
		},

		{
			name:      "from shift out to information separator",
			faultsLen: 1,
			chars: []string{"\u000e", "\u000f", "\u0010", "\u0011", "\u0012",
				"\u0013", "\u0014", "\u0015", "\u0016", "\u0017", "\u0018",
				"\u0019", "\u001a", "\u001b", "\u001c", "\u001d", "\u001e", "\u001f"},
		},

		{
			name:      "from delete to control",
			faultsLen: 1,
			chars: []string{"\u007f", "\u0080", "\u0081", "\u0082",
				"\u0083", "\u0084", "\u0085", "\u0086", "\u0087",
				"\u0088", "\u0089", "\u008a", "\u008b", "\u008c", "\u008d",
				"\u008e", "\u008f", "\u0090", "\u0091", "\u0092",
				"\u0093", "\u0094", "\u0095", "\u0096", "\u0097",
				"\u0098", "\u0099", "\u009a", "\u009b", "\u009c", "\u009d",
				"\u009e", "\u009f"},
		},

		{
			name:      "soft hyphen",
			faultsLen: 1,
			chars:     []string{"\u00ad"},
		},

		{
			name:      "Arabic number sign range",
			faultsLen: 1,
			chars:     []string{"\u0600", "\u0601", "\u0602", "\u0603", "\u0604"},
		},

		{
			name:      "Syriac Abbreviation, Khmer",
			faultsLen: 1,
			chars:     []string{"\u070f", "\u17b4", "\u17b5"},
		},

		{
			name:      "Zero width non-joiner range",
			faultsLen: 1,
			chars:     []string{"\u200c", "\u200d", "\u200e", "\u200f"},
		},

		{
			name:      "from line separator to narrow no-break space",
			faultsLen: 1,
			chars: []string{"\u2028", "\u2029", "\u202a", "\u202b",
				"\u202c", "\u202d", "\u202e", "\u202f"},
		},

		{
			name:      "from word joiner to nominal digit shapes",
			faultsLen: 1,
			chars: []string{"\u2060", "\u2061", "\u2062", "\u2063",
				"\u2064", "\u2065", "\u2066", "\u2067",
				"\u2068", "\u2069", "\u206a", "\u206b",
				"\u206c", "\u206d", "\u206e", "\u206f"},
		},

		{
			name:      "zero width no-break space",
			faultsLen: 1,
			chars:     []string{"\ufeff"},
		},

		{
			name:      "not valid unicode characters range",
			faultsLen: 1,
			chars: []string{"\ufff0", "\ufff1", "\ufff2", "\ufff3",
				"\ufff4", "\ufff5", "\ufff6", "\ufff7",
				"\ufff8", "\ufff9", "\ufffa", "\ufffb",
				"\ufffc", "\ufffd", "\ufffe", "\uffff"},
		},
	}

	for _, tcase := range tcases {
		for _, ch := range tcase.chars {
			page := bytes.NewBufferString(fmt.Sprintf("<a attr='%s'/>", ch))
			faults, err := reviewer.Review("something.html", page)

			r.NoError(err, tcase.name)
			r.Len(faults, tcase.faultsLen, tcase.name)
			if tcase.faultsLen == 0 {
				continue
			}

			r.Equal(faults[0].Rule.Code, reviewers.Rules[reviewer.ReviewerName()].Code)
			r.Equal(faults[0].Rule.Description, reviewers.Rules[reviewer.ReviewerName()].Description)
			r.Equal(faults[0].Reviewer, reviewer.ReviewerName())
			r.Equal("something.html", faults[0].Path)
		}
	}

}
