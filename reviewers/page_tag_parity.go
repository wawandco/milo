package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

var selfClosing = []string{
	"area",
	"base",
	"br",
	"col",
	"command",
	"embed",
	"hr",
	"img",
	"input",
	"keygen",
	"link",
	"menuitem",
	"meta",
	"param",
	"source",
	"track",
	"wbr",
}

// PageTagParity is a reviewer that checks that all img tags have alt attribute.
type PageTagParity struct{}

// ReviewerName returns the reviewer name.
func (t PageTagParity) ReviewerName() string {
	return "page-tag-parity"
}

// Accepts checks if the file can be reviewed.
func (t PageTagParity) Accepts(path string) bool {
	return true
}

// Review returns a fault for each tag that doesn't have its end tag.
// For example:
// <div><div>...</div></div> is correct.
// <div><div>...</div> is a fault.
//
// Self-closed tags will not generate any fault.
// For example: <br/>, <input/>, <img/>.
func (t PageTagParity) Review(path string, page io.Reader) ([]Fault, error) {
	var fault []Fault
	var openedTags []*html.Token
	var err error

	z := html.NewTokenizer(page)
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
			if t.isSelfClosing(token) {
				continue
			}

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
					Col:      token.Col,
					Path:     path,
					Rule:     Rules[t.ReviewerName()],
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
				Col:      o.Col,
				Path:     path,
				Rule:     Rules[t.ReviewerName()],
				Reviewer: t.ReviewerName(),
			})
		}
	}

	return fault, nil
}

func (t PageTagParity) isSelfClosing(token html.Token) bool {
	for _, selfc := range selfClosing {
		if selfc != token.DataAtom.String() {
			continue
		}

		return true
	}

	return false
}
