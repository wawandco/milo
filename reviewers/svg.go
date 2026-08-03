package reviewers

import (
	"github.com/wawandco/milo/internal/html"
)

// svgTracker tracks whether the tokenizer is currently inside an <svg> element.
//
// SVG content follows SVG/XML conventions rather than HTML ones: camelCase
// attributes (viewBox, strokeWidth), namespaced attributes (xlink:href) and
// self-closed elements (<path>) that are not HTML void elements. Reviewers
// use this tracker to skip that content instead of flagging it with HTML rules.
type svgTracker struct {
	depth int
}

// inSVG updates the tracker with the current token and reports whether the
// token belongs to the SVG space, including the boundary <svg> tag itself,
// whose attributes (e.g. viewBox) already follow SVG conventions.
func (s *svgTracker) inSVG(tt html.TokenType, tok html.Token) bool {
	isSVGTag := tok.Data == "svg"

	switch tt {
	case html.StartTagToken:
		if isSVGTag {
			s.depth++
		}
	case html.SelfClosingTagToken:
		return s.depth > 0 || isSVGTag
	case html.EndTagToken:
		inside := s.depth > 0
		if isSVGTag && s.depth > 0 {
			s.depth--
		}

		return inside
	}

	return s.depth > 0
}

// inSVGContent works like inSVG but reports false for the boundary <svg> and
// </svg> tags. Reviewers checking HTML tag structure (e.g. tag parity) use it
// so the <svg> element itself is still verified while its content is skipped.
func (s *svgTracker) inSVGContent(tt html.TokenType, tok html.Token) bool {
	isSVGTag := tok.Data == "svg"
	inside := s.depth > 0

	switch tt {
	case html.StartTagToken:
		if isSVGTag {
			s.depth++

			return false
		}
	case html.SelfClosingTagToken:
		return inside || isSVGTag
	case html.EndTagToken:
		if isSVGTag && s.depth > 0 {
			s.depth--

			return false
		}
	}

	return inside
}
