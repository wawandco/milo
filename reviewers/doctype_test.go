package reviewers_test

import (
	"strings"
	"testing"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_DoctypeReviewer_Review(t *testing.T) {
	r := require.New(t)

	doc := reviewers.Doctype{}
	page := strings.NewReader("<html></html>")

	faults, err := doc.Review(page)
	r.NoError(err)

	r.Len(faults, 1)
	r.Equal(faults[0].ReviewerName, doc.ReviewerName())
}
