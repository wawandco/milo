package reviewers

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type AttrNoDuplication struct{}

func (a AttrNoDuplication) ReviewerName() string {
	return "attribute/no-duplication"
}

func (a AttrNoDuplication) Accepts(path string) bool {
	return true
}

func (a AttrNoDuplication) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tok := z.Token()
		attrNames := a.inlineAttributesName(tok.Attr)
		for _, attr := range tok.Attr {
			if strings.Count(attrNames, attr.Key) > 1 {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Path:     path,
					Rule:     Rules["0010"],
				})
			}
			attrNames = strings.ReplaceAll(attrNames, attr.Key, "")
		}
	}

	return fault, nil
}

func (a AttrNoDuplication) inlineAttributesName(attributes []html.Attribute) string {
	var attrNames []string
	for _, atr := range attributes {
		attrNames = append(attrNames, atr.Key)
	}
	return strings.Join(attrNames, ",")
}
