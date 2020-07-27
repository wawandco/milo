package reviewers

import (
	"io"

	"github.com/wawandco/milo/internal/html"
)

type TagPair struct{}

func (t TagPair) ReviewerName() string {
	return "tag/pair"
}

func (t TagPair) Accepts(path string) bool {
	return true
}

func (t TagPair) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault
	var openedTags []*html.Token
	var err error

	z := html.NewTokenizer(r)
	for {
		tt := z.Next()

		if err = z.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return []Fault{}, err
		}

		token := z.Token()
		switch tt {
		case html.StartTagToken:
			openedTags = append(openedTags, &token)

		case html.EndTagToken:

			var i int
			var levels int
			for i = len(openedTags) - 1; i >= 0; i-- {
				if openedTags[i] == nil {
					continue
				}
				if openedTags[i].DataAtom == token.DataAtom {
					openedTags[i] = nil
					break
				}
				if openedTags[i].DataAtom != 0 {
					levels++
				}
			}

			if i == -1 {
				fault = append(fault, Fault{
					Line:     token.Line,
					Path:     path,
					Rule:     Rules["0015"],
					Reviewer: t.ReviewerName(),
				})
				continue
			}

			// Mark all open tags as consumed after a tag is matched if spaces > 0.
			// it means we skipped some single open tags.
			if levels > 0 {
				for i = 0; i < len(openedTags); i++ {
					if openedTags[i] != nil {
						openedTags[i].DataAtom = 0
					}
				}
			}
		}
	}

	for _, o := range openedTags {
		if o != nil {
			fault = append(fault, Fault{
				Line:     o.Line,
				Path:     path,
				Rule:     Rules["0015"],
				Reviewer: t.ReviewerName(),
			})
		}
	}

	return fault, nil
}
