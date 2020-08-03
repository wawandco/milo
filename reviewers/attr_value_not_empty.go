package reviewers

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/wawandco/milo/external/html"
)

var (
	rgx = regexp.MustCompile(`\<\%\=?[^\>]+?\%\>`)
)

type AttrValueNotEmpty struct{}

func (a AttrValueNotEmpty) ReviewerName() string {
	return "attribute/value-not-empty"
}

func (a AttrValueNotEmpty) Accepts(path string) bool {
	return true
}

func (a AttrValueNotEmpty) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if rgx.Match(z.Raw()) {
			b := rgx.ReplaceAll(z.Raw(), []byte{})
			b = append(b, z.Buffered()...)
			z = html.NewTokenizer(bytes.NewBuffer(b))
		}

		tok := z.Token()
		for _, attr := range tok.Attr {
			if strings.TrimSpace(attr.Val) == "" {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Path:     path,
					Rule:     Rules["0011"],
				})
			}
		}
	}

	return fault, nil
}
