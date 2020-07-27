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
		Name:        "attribute/style",
		Description: "inline-style-disabled: Inline style cannot be used",
	},

	"0010": {
		Code:        "0010",
		Name:        "attribute/no-duplication",
		Description: "attr-no-duplication: Elements cannot have duplicate attributes",
	},

	"0011": {
		Code:        "0011",
		Name:        "attribute/value-not-empty",
		Description: "attr-value-not-empty: All attributes must have values",
	},

	"0012": {
		Code:        "0012",
		Name:        "attribute/alt-required",
		Description: "alt-required: Alt attribute required for images, areas and input[type=image]",
	},

	"0013": {
		Code:        "0012",
		Name:        "attribute/lowercase",
		Description: "attr-lowercase: attributes should be in lowercase",
	},

	"0014": {
		Code:        "0014",
		Name:        "attribute/id-unique",
		Description: "attr-id-unique: id attribute must be unique",
	},

	"0015": {
		Code:        "0015",
		Name:        "tag/pair",
		Description: "tag-pair: Tag must be paired",
	},

	"0016": {
		Code:        "0016",
		Name:        "attribute/unsafe-chars",
		Description: "attr-unsafe-chars: Attribute values cannot contain unsafe chars",
	},

	"0017": {
		Code:        "0017",
		Name:        "script/inline-disabled",
		Description: "inline-script-disabled: Inline script cannot be used",
	},

	"0018": {
		Code:        "0018",
		Name:        "attribute/double-quotes",
		Description: "attr-value-double-quotes: Attribute values must be in double quotes",
	},

	"0019": {
		Code:        "0018",
		Name:        "attribute/no-white-spaces",
		Description: "attr-no-white-spaces: should not have spaces",
	},
}
