package reviewers_test

import (
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func Test_StyleTag_Review(t *testing.T) {
	reviewer := reviewers.PageStyleTagForbidden{}
	tcases := []struct {
		name      string
		content   string
		err       error
		faultsLen int
		faults    []reviewers.Fault
	}{
		{
			name:      "no inline css",
			faultsLen: 0,
			content:   "<div></div>",
		},

		{
			name:      "style tag present in partial",
			faultsLen: 1,
			content:   `<div> <style class=""></style></div>`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     1,
					Col:      7,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "style tag present full page",
			faultsLen: 1,
			content: `
			<html>
				<head><head>
				<body
					<div> <STYLE></STYLE></div>
				<body>
			<html>
			`,
			faults: []reviewers.Fault{
				{
					Reviewer: reviewer.ReviewerName(),
					Line:     5,
					Col:      12,
					Rule:     reviewers.Rules[reviewer.ReviewerName()],
				},
			},
		},

		{
			name:      "style tag present in comment",
			faultsLen: 0,
			content: `
			<html>
				<head><head>
				<body
					<div> <!-- <STYLE></STYLE> --></div>
				<body>
			<html>
			`,
		},
	}

	for _, tcase := range tcases {
		page := strings.NewReader(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(faults) != tcase.faultsLen {
			t.Errorf("expected %v, got %v", tcase.faultsLen, len(faults))
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
			if "something.html" != faults[index].Path {
				t.Errorf("expected %v, got %v", "something.html", faults[index].Path)
			}
		}

	}

}

func Test_StyleTag_Accept(t *testing.T) {
	doc := reviewers.PageStyleTagForbidden{}

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
