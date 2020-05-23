package milo_test

import (
	"strings"
	"testing"
	"wawandco/milo"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_Referee(t *testing.T) {
	r := require.New(t)
	referee := milo.NewReferee()

	reader := strings.NewReader("<html></html>")
	faults, err := referee.Review("something.html", reader)

	r.NoError(err)
	r.Len(faults, 0)

	referee.Reviewers = []milo.Reviewer{
		reviewers.DoctypePresent{},
	}

	faults, err = referee.Review("something.html", reader)

	r.NoError(err)
	r.Len(faults, 1)

}
