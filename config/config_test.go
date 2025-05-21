package config

import (
	_ "embed"
	"os"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

//go:embed testdata/config.yaml
var configTemplate []byte

func Test_Load(t *testing.T) {
	t.Run("config file is present", func(t *testing.T) {
		wd := t.TempDir()
		cwd, _ := os.Getwd()
		t.Cleanup(func() {
			os.Chdir(cwd)
			os.RemoveAll(wd)
		})

		if err := os.Chdir(wd); err != nil {
			t.Fatalf("error changing directory: %v", err)
		}
		
		if err := os.WriteFile(".milo.yml", configTemplate, 0777); err != nil {
			t.Fatalf("error writing config file: %v", err)
		}

		config, err := Load()
		if err != nil {
			t.Fatalf("error loading config: %v", err)
		}
		
		if got, want := len(config.Reviewers), 1; got != want {
			t.Errorf("expected %d reviewers, got %d", want, got)
		}
		
		if got, want := len(config.SelectedReviewers()), 1; got != want {
			t.Errorf("expected %d selected reviewers, got %d", want, got)
		}
	})

	t.Run("no config file", func(t *testing.T) {
		if err := os.Chdir(t.TempDir()); err != nil {
			t.Fatalf("error changing directory: %v", err)
		}

		config, err := Load()
		if err != nil {
			t.Fatalf("error loading config: %v", err)
		}
		
		if got, want := len(config.Reviewers), 0; got != want {
			t.Errorf("expected %d reviewers, got %d", want, got)
		}
		
		if got, want := len(config.SelectedReviewers()), len(reviewers.All); got != want {
			t.Errorf("expected %d selected reviewers, got %d", want, got)
		}
	})
}
