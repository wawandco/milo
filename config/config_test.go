package config

import (
	_ "embed"
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/reviewers"
)

//go:embed testdata/config.yaml
var configTemplate []byte

func Test_Load(t *testing.T) {
	t.Run("config file is present", func(t *testing.T) {
		r := is.New(t)

		wd := t.TempDir()
		cwd, _ := os.Getwd()
		t.Cleanup(func() {
			os.Chdir(cwd)
			os.RemoveAll(wd)
		})

		r.NoErr(os.Chdir(wd))
		r.NoErr(os.WriteFile(".milo.yml", configTemplate, 0777))

		config, err := Load()
		r.NoErr(err)
		r.Equal(len(config.Reviewers), 1)
		r.Equal(len(config.SelectedReviewers()), 1)
	})

	t.Run("no config file", func(t *testing.T) {
		r := is.New(t)
		r.NoErr(os.Chdir(t.TempDir()))

		config, err := Load()
		r.NoErr(err)
		r.Equal(len(config.Reviewers), 0)
		r.Equal(len(config.SelectedReviewers()), len(reviewers.All))
	})
}
