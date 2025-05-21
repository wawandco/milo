package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_TagLowercase_Review(t *testing.T) {
	reviewer := reviewers.PageTagLowercaseRequired{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:    "no inline css",
			content: "<div></div>",
		},

		{
			name:    "lowercased",
			content: `<div><style class=""></style></div>`,
		},

		{
			name: "UPPER case tag",
			content: `
			<html>
				<head><head>
				<body>
					<div><STYLE></STYLE></div>
				<body>
			<html>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      18,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name: "mixed cases on tag name",
			content: `
			<html>
				<Head><Head>
				<body
					<div><STYLE></STYLE></div>
				<body>
			<html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      5,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     3,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      11,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      18,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(faults) != len(tcase.faults) {
			t.Errorf("expected length %d, got %d", len(tcase.faults), len(faults))
		}

		if tcase.faultsLen == 0 {
			continue
		}

		for index, fault := range tcase.faults {
			if fault.Reviewer != faults[index].Reviewer {
				t.Errorf("expected %v, got %v", fault.Reviewer, faults[index].Reviewer)
			}
			if fault.Line != faults[index].Line {
				t.Errorf("expected %v, got %v", fault.Line, faults[index].Line)
			}
			if fault.Col != faults[index].Col {
				t.Errorf("expected %v, got %v", fault.Col, faults[index].Col)
			}
			if fault.Rule.Code != faults[index].Rule.Code {
				t.Errorf("expected %v, got %v", fault.Rule.Code, faults[index].Rule.Code)
			}
			if fault.Rule.Description != faults[index].Rule.Description {
				t.Errorf("expected %v, got %v", fault.Rule.Description, faults[index].Rule.Description)
			}
			if "something.html" != fault.Path {
				t.Errorf("expected %v, got %v", "something.html", fault.Path)
			}
		}

	}

}

func Test_TagLowercase_Accept(t *testing.T) {
	doc := reviewers.PageTagLowercaseRequired{}

	if !doc.Accepts("/very/long/path/name/_partial.plush.html") {
		t.Error("Expected to accept /very/long/path/name/_partial.plush.html")
	}
	if !doc.Accepts("_partial.plush.html") {
		t.Error("Expected to accept _partial.plush.html")
	}
	if !doc.Accepts("page.plush.html") {
		t.Error("Expected to accept page.plush.html")
	}
	if !doc.Accepts("templates/_partial.plush.html") {
		t.Error("Expected to accept templates/_partial.plush.html")
	}
}

func Test_TagLowercase_Name(t *testing.T) {
	doc := reviewers.PageTagLowercaseRequired{}
	if doc.ReviewerName() != "page-tag-lowercase" {
		t.Errorf("expected %v, got %v", "page-tag-lowercase", doc.ReviewerName())
	}
}
