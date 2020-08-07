package reviewers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_InlineScriptDisabled_Review(t *testing.T) {
	r := require.New(t)

	reviewer := reviewers.InlineScriptDisabled{}
	tcases := []struct {
		name      string
		events    []string
		faultsLen int
	}{
		{
			name:      "onttt should not result in an error",
			faultsLen: 0,
			events:    []string{"onttt"},
		},

		{
			name:      "on load/on unload",
			faultsLen: 1,
			events: []string{"onload", "onunload", "onLoad",
				"onUnload", "OnLoAd", "OnUnLoAd"},
		},

		{
			name:      "on message",
			faultsLen: 1,
			events:    []string{"onmessage", "onMessage", "OnMessage"},
		},

		{
			name:      "on error",
			faultsLen: 1,
			events:    []string{"onerror", "onError", "OnErRoR"},
		},

		{
			name:      "on submit",
			faultsLen: 1,
			events:    []string{"onsubmit", "onSubmit", "OnSubmit"},
		},

		{
			name:      "on select",
			faultsLen: 1,
			events:    []string{"onselect", "onSelect", "OnSelect"},
		},

		{
			name:      "on change",
			faultsLen: 1,
			events:    []string{"onchange", "onChange", "OnChAnGe"},
		},

		{
			name:      "on sroll",
			faultsLen: 1,
			events:    []string{"onscroll", "onScroll", "OnScroll"},
		},

		{
			name:      "on resize",
			faultsLen: 1,
			events:    []string{"onresize", "onResize", "OnResize"},
		},

		{
			name:      "on mouse events",
			faultsLen: 1,
			events: []string{"onmouseover", "onmouseout", "onmousemove",
				"onmouseleave", "onmouseenter", "onmousedown",
			},
		},

		{
			name:      "on key events",
			faultsLen: 1,
			events:    []string{"onkeyup", "onkeypress", "onkeydown"},
		},

		{
			name:      "on focus and blur",
			faultsLen: 1,
			events:    []string{"onfocus", "onblur"},
		},

		{
			name:      "on click and double click",
			faultsLen: 1,
			events:    []string{"onclick", "ondblclick"},
		},

		{
			name:      "javascript protocol [ javascript: ] should result in error for src",
			faultsLen: 1,
			events:    []string{`src="javascript:alert(1)"`, `src="   JAVASCRIPT:alert(2)"`},
		},

		{
			name:      "javascript protocol [ javascript: ] should result in error for href",
			faultsLen: 1,
			events:    []string{`href="javascript:alert(1)"`, `href="   JAVASCRIPT:alert(2)"`},
		},
	}

	for _, tcase := range tcases {
		for _, evt := range tcase.events {
			page := bytes.NewBufferString(fmt.Sprintf("<a %s />", evt))
			faults, err := reviewer.Review("something.html", page)

			r.NoError(err, tcase.name)
			r.Len(faults, tcase.faultsLen, tcase.name)
			if tcase.faultsLen == 0 {
				continue
			}

			r.Equal(faults[0].Rule.Code, reviewers.Rules[reviewer.ReviewerName()].Code)
			r.Equal(faults[0].Rule.Description, reviewers.Rules[reviewer.ReviewerName()].Description)
			r.Equal(faults[0].Reviewer, reviewer.ReviewerName())
			r.Equal("something.html", faults[0].Path)
		}
	}

}
