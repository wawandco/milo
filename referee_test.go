package milo_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"wawandco/milo"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_Referee(t *testing.T) {
	r := require.New(t)
	referee := milo.NewReferee()

	path := filepath.Join(os.TempDir(), "something.html")
	reader := strings.NewReader("<html></html>")

	faults, err := referee.Review(path, reader)

	r.NoError(err)
	r.Len(faults, 0)

	referee.Reviewers = []milo.Reviewer{
		reviewers.DoctypePresent{},
	}

	reader = strings.NewReader("<html></html>")
	faults, err = referee.Review(path, reader)

	r.NoError(err)
	r.Len(faults, 1)
}

func Test_RefereeMultiple(t *testing.T) {
	r := require.New(t)
	referee := milo.NewReferee()
	referee.Reviewers = []milo.Reviewer{
		reviewers.DoctypePresent{},
		reviewers.DoctypeValid{},
	}

	reader := strings.NewReader(`
		<!DOCTYPE invalid>
		<html></html>
	`)

	faults, err := referee.Review("something.html", reader)
	r.NoError(err)
	r.Len(faults, 1)
}
