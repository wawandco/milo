package reviewers

type Rule struct {
	Code        string
	Name        string
	Description string
}

// Rules catalog for usage.
var Rules = map[string]Rule{
	"0001": {
		Code:        "0001",
		Name:        "doctype/present",
		Description: "doctype tag must be present in the document",
	},

	"0002": {
		Code:        "0002",
		Name:        "doctype/valid",
		Description: "doctype tag must be valid",
	},

	"0003": {
		Code:        "0003",
		Name:        "css/inline",
		Description: "don't use inline css for other reasons than quick testing locally",
	},

	"0004": {
		Code:        "0004",
		Name:        "title/present",
		Description: "<title> tag should be present and have content inside",
	},

	"0005": {
		Code:        "0005",
		Name:        "style/tag-present",
		Description: "<style> tag should not be used",
	},

	"0006": {
		Code:        "0006",
		Name:        "tag/lowercase",
		Description: "tag names must be in lowercase",
	},

	"0007": {
		Code:        "0007",
		Name:        "tag/src-empty",
		Description: "`src`, `href` and `data` attributes of must have a value",
	},

	"0008": {
		Code:        "0008",
		Name:        "tag/ul-ol-li",
		Description: "UL and OL must only have LI direct children",
	},

	"0009": {
		Code:        "0009",
		Name:        "tag/ul-ol-li",
		Description: "UL and OL must only have LI direct children",
	},
}
