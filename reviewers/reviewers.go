package reviewers

// Rules catalog for usage.
var Rules = map[string]Rule{
	"0001": Rule{
		Code:        "0001",
		Name:        "doctype/present",
		Description: "doctype tag must be present in the document",
	},

	"0002": Rule{
		Code:        "0002",
		Name:        "doctype/valid",
		Description: "doctype tag must be valid",
	},

	"0003": Rule{
		Code:        "0003",
		Name:        "css/inline",
		Description: "don't use inline css for other reasons than quick testing locally",
	},
}
