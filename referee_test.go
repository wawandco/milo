package milo_test

import (
	"strings"
	"testing"
	"wawandco/milo"

	"github.com/stretchr/testify/require"
)

func Test_Referee(t *testing.T) {
	r := require.New(t)
	referee := milo.NewReferee()
	referee.Reviewers = []milo.Reviewer{}

	reader := strings.NewReader("<html></html>")
	faults, err := referee.Review(reader)

	r.NoError(err)
	r.Len(faults, 0)
}
