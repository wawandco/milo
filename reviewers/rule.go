package reviewers

type Rule struct {
	Code        string
	Name        string
	Description string
}

// Rules catalog for usage.
var Rules = map[string]Rule{
	"doctype/present": {
		Code:        "0001",
		Name:        "doctype/present",
		Description: "doctype tag must be present in the document",
	},

	"doctype/valid": {
		Code:        "0002",
		Name:        "doctype/valid",
		Description: "doctype tag must be valid",
	},

	"css/inline": {
		Code:        "0003",
		Name:        "css/inline",
		Description: "don't use inline css for other reasons than quick testing locally",
	},

	"title/present": {
		Code:        "0004",
		Name:        "title/present",
		Description: "<title> tag should be present and have content inside",
	},

	"style/tag-present": {
		Code:        "0005",
		Name:        "style/tag-present",
		Description: "<style> tag should not be used",
	},

	"tag/lowercase": {
		Code:        "0006",
		Name:        "tag/lowercase",
		Description: "tag names must be in lowercase",
	},

	"tag/src-empty": {
		Code:        "0007",
		Name:        "tag/src-empty",
		Description: "`src`, `href` and `data` attributes of must have a value",
	},

	"tag/ul-ol-li": {
		Code:        "0008",
		Name:        "tag/ul-ol-li",
		Description: "UL and OL must only have LI direct children",
	},

	"attribute/style": {
		Code:        "0009",
		Name:        "attribute/style",
		Description: "inline-style-disabled: Inline style cannot be used",
	},

	"attribute/no-duplication": {
		Code:        "0010",
		Name:        "attribute/no-duplication",
		Description: "attr-no-duplication: Elements cannot have duplicate attributes",
	},

	"attribute/value-not-empty": {
		Code:        "0011",
		Name:        "attribute/value-not-empty",
		Description: "attr-value-not-empty: All attributes must have values",
	},

	"attribute/alt-required": {
		Code:        "0012",
		Name:        "attribute/alt-required",
		Description: "alt-required: Alt attribute required for images, areas and input[type=image]",
	},

	"attribute/lowercase": {
		Code:        "0012",
		Name:        "attribute/lowercase",
		Description: "attr-lowercase: attributes should be in lowercase",
	},

	"attribute/id-unique": {
		Code:        "0014",
		Name:        "attribute/id-unique",
		Description: "attr-id-unique: id attribute must be unique",
	},

	"tag/pair": {
		Code:        "0015",
		Name:        "tag/pair",
		Description: "tag-pair: Tag must be paired",
	},

	"attribute/unsafe-chars": {
		Code:        "0016",
		Name:        "attribute/unsafe-chars",
		Description: "attr-unsafe-chars: Attribute values cannot contain unsafe chars",
	},

	"script/inline-disabled": {
		Code:        "0017",
		Name:        "script/inline-disabled",
		Description: "inline-script-disabled: Inline script cannot be used",
	},

	"attribute/double-quotes": {
		Code:        "0018",
		Name:        "attribute/double-quotes",
		Description: "attr-value-double-quotes: Attribute values must be in double quotes",
	},

	"attribute/no-white-spaces": {
		Code:        "0019",
		Name:        "attribute/no-white-spaces",
		Description: "attr-no-white-spaces: should be whitespaces between the attribute and value",
	},
}
