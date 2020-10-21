package reviewers

import (
	"io"
)

// All the reviewers we have build.
var All = []Reviewer{
	PageDoctypePresent{},
	PageDoctypeValid{},
	PageInlineCSSForbidden{},
	PageTitlePresent{},
	PageStyleTagForbidden{},
	PageTagLowercaseRequired{},
	AttributeSrcRequired{},
	TagOlUlOnlyLiChild{},
	AttributeNoDuplication{},
	AttributeValueNotEmpty{},
	AttributeLowercase{},
	AttributeIDUnique{},
	AttributeAltRequired{},
	PageTagParity{},
	AttributeUnsafeChars{},
	PageInlineScriptForbidden{},
	AttributeValueDoubleQuotes{},
}

// Reviewer is in charge of reviewing a path and a content and return the list of faults on it.
// for each reviewer there is typically a rule that it will check.
type Reviewer interface {
	ReviewerName() string
	Accepts(fileName string) bool
	Review(path string, content io.Reader) ([]Fault, error)
}
