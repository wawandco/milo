package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/external/html"
)

type SrcEmpty struct{}

func (css SrcEmpty) ReviewerName() string {
	return "tag/src-empty"
}

func (css SrcEmpty) Accepts(path string) bool {
	return true
}

func (css SrcEmpty) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	z := html.NewTokenizer(page)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return []Fault{}, err
		}

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		for _, attr := range tok.Attr {
			if (attr.Key == "src" || attr.Key == "href" || attr.Key == "data") &&
				strings.TrimSpace(attr.Val) == "" {
				result = append(result, Fault{
					Reviewer: css.ReviewerName(),
					Line:     attr.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[css.ReviewerName()],
				})
			}
		}
	}

	return result, nil
}
