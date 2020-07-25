package reviewers

import (
	"io"
)

// All the reviewers we have build.
var All = []Reviewer{
	DoctypePresent{},
	DoctypeValid{},
	InlineCSS{},
	TitlePresent{},
	StyleTag{},
	TagLowercase{},
	SrcEmpty{},
	OlUlValid{},
	AttrNoDuplication{},
	AttrValueNotEmpty{},
	AttrLowercase{},
	AttrIDUnique{},
	AltRequired{},
	TagPair{},
	AttrUnsafeChars{},
	InlineScriptDisabled{},
}

type Reviewer interface {
	ReviewerName() string
	Accepts(fileName string) bool
	Review(path string, content io.Reader) ([]Fault, error)
}
