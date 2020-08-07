package reviewers

import (
	"io"
	"regexp"

	"github.com/wawandco/milo/external/html"
)

type AttrUnsafeChars struct{}

func (a AttrUnsafeChars) ReviewerName() string {
	return "attribute/unsafe-chars"
}

func (a AttrUnsafeChars) Accepts(path string) bool {
	return true
}

func (a AttrUnsafeChars) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault
	var unsafeChars = regexp.MustCompile("[\u0000-\u0009\u000b\u000c\u000e-\u001f\u007f-\u009f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]")

	z := html.NewTokenizer(r)
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
					Path:     path,
					Rule:     Rules[a.ReviewerName()],
				})
			}
		}
	}

	return fault, nil
}
