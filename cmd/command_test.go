package cmd

import (
	"errors"
	"testing"

	"github.com/wawandco/milo/internal/assert"
)

// TestRunner is a mock implementation of the Runner interface for testing
type TestRunner struct {
	NameValue  string
	RunFunc    func([]string) error
	RunCalled  bool
	RunArgs    []string
}

func (r *TestRunner) Name() string {
	return r.NameValue
}

func (r *TestRunner) Run(args []string) error {
	r.RunCalled = true
	r.RunArgs = args
	if r.RunFunc != nil {
		return r.RunFunc(args)
	}
	return nil
}

// TestHelpProvider is a mock implementation of the HelpProvider interface for testing
type TestHelpProvider struct {
	NameValue      string
	HelpTextValue  string
}

func (h *TestHelpProvider) Name() string {
	return h.NameValue
}

func (h *TestHelpProvider) HelpText() string {
	return h.HelpTextValue
}

func TestRunnerInterface(t *testing.T) {
	// Test successful command run
	t.Run("successful run", func(t *testing.T) {
		runner := &TestRunner{
			NameValue: "test-command",
			RunFunc: func(args []string) error {
				return nil
			},
		}

		args := []string{"--flag", "value"}
		err := runner.Run(args)

		assert.NoError(t, err)
		assert.True(t, runner.RunCalled)
		assert.Equal(t, args, runner.RunArgs)
		assert.Equal(t, "test-command", runner.Name())
	})

	// Test command run with error
	t.Run("run with error", func(t *testing.T) {
		expectedErr := errors.New("command failed")
		runner := &TestRunner{
			NameValue: "failing-command",
			RunFunc: func(args []string) error {
				return expectedErr
			},
		}

		err := runner.Run(nil)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, "failing-command", runner.Name())
	})
}

func TestHelpProviderInterface(t *testing.T) {
	t.Run("help provider", func(t *testing.T) {
		provider := &TestHelpProvider{
			NameValue:     "help-command",
			HelpTextValue: "This is help text for testing",
		}

		assert.Equal(t, "help-command", provider.Name())
		assert.Equal(t, "This is help text for testing", provider.HelpText())
		assert.NotEqual(t, "wrong text", provider.HelpText())
	})
}