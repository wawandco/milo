package milo_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"wawandco/milo"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

func Test_Referee(t *testing.T) {
	r := require.New(t)
	referee := milo.NewReferee()

	path := filepath.Join(os.TempDir(), "something.html")
	err := ioutil.WriteFile(path, []byte("<html></html>"), 0644)
	r.NoError(err)

	faults, err := referee.Review(path)

	r.NoError(err)
	r.Len(faults, 0)

	referee.Reviewers = []milo.Reviewer{
		reviewers.DoctypePresent{},
	}

	faults, err = referee.Review(path)

	r.NoError(err)
	r.Len(faults, 1)

}
