package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
)

var configTemplate = `
output: "github"
reviewers:
  - doctype-present
`

func Test_Load(t *testing.T) {
	r := require.New(t)
	wd := os.TempDir()

	cwd, err := os.Getwd()
	r.NoError(err)
	defer func() {
		os.Chdir(cwd)
		os.RemoveAll(wd)
	}()

	err = os.WriteFile(filepath.Join(wd, ".milo.yml"), []byte(configTemplate), 0777)
	r.NoError(err)

	os.Chdir(wd)

	config, err := Load()
	r.NoError(err)
	r.Len(config.Reviewers, 1)
	r.Len(config.SelectedReviewers(), 1)

	err = os.Remove(filepath.Join(wd, ".milo.yml"))
	r.NoError(err)

	config, err = Load()
	r.NoError(err)
	r.Len(config.Reviewers, 0)
	r.Len(config.SelectedReviewers(), len(reviewers.All))
}
