package reviewers_test

import (
	"bytes"
	"testing"

	"github.com/wawandco/milo/internal/assert"
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
		page := bytes.NewBufferString(tcase.content)
		faults, err := reviewer.Review("something.html", page)

		assert.NoError(t, err)
		assert.Equal(t, tcase.faultsLen, len(faults))

		if tcase.faultsLen == 0 {
			continue
		}

		// Verify each fault matches the expected values
		assert.Faults(t, faults, tcase.faults)

	}

}

func Test_StyleTag_Accept(t *testing.T) {
	doc := reviewers.PageStyleTagForbidden{}

	assert.True(t, doc.Accepts("/very/long/path/name/_partial.plush.html"), 
		"Expected to accept /very/long/path/name/_partial.plush.html")
	assert.True(t, doc.Accepts("_partial.plush.html"),
		"Expected to accept _partial.plush.html")
	assert.True(t, doc.Accepts("page.plush.html"),
		"Expected to accept page.plush.html")
	assert.True(t, doc.Accepts("templates/_partial.plush.html"),
		"Expected to accept templates/_partial.plush.html")
}
