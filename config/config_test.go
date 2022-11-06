package config

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wawandco/milo/reviewers"
)

//go:embed testdata/config.yaml
var configTemplate []byte

func Test_Load(t *testing.T) {
	t.Run("config file is present", func(t *testing.T) {
		r := require.New(t)

		wd := t.TempDir()
		cwd, _ := os.Getwd()
		t.Cleanup(func() {
			os.Chdir(cwd)
			os.RemoveAll(wd)
		})

		r.NoError(os.Chdir(wd))
		r.NoError(os.WriteFile(".milo.yml", configTemplate, 0777))

		config, err := Load()
		r.NoError(err)
		r.Len(config.Reviewers, 1)
		r.Len(config.SelectedReviewers(), 1)
	})

	t.Run("no config file", func(t *testing.T) {
		r := require.New(t)
		r.NoError(os.Chdir(t.TempDir()))

		config, err := Load()
		r.NoError(err)
		r.Len(config.Reviewers, 0)
		r.Len(config.SelectedReviewers(), len(reviewers.All))
	})
}
