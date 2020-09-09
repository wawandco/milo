package reviewers

import (
	"io"
	"regexp"

	"github.com/wawandco/milo/external/html"
)

// AttrUnsafeChars is a reviewer that checks that tags does not have unsafe characters into teir attribute values.
type AttrUnsafeChars struct{}

// ReviewerName returns the reviewer name.
func (a AttrUnsafeChars) ReviewerName() string {
	return "attribute/unsafe-chars"
}

// Accepts checks if the file can be reviewed.
func (a AttrUnsafeChars) Accepts(path string) bool {
	return true
}

// Review return a fault for each attribute value that contains an unsage character.
func (a AttrUnsafeChars) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault
	var unsafeChars = regexp.MustCompile("[\u0000-\u0009\u000b\u000c\u000e-\u001f\u007f-\u009f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]")

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tok := z.Token()
		for _, attr := range tok.Attr {
			if unsafeChars.MatchString(attr.Val) {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[a.ReviewerName()],
				})
			}
		}
	}

	return fault, nil
}
