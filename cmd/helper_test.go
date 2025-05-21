package cmd

import (
	"fmt"
	"testing"

	"github.com/wawandco/milo/internal/assert"
)

// TestHelpProvider is a mock implementation of the HelpProvider interface for testing
type testHelpProvider struct {
	nameValue     string
	helpTextValue string
}

func (h *testHelpProvider) Name() string {
	return h.nameValue
}

func (h *testHelpProvider) HelpText() string {
	return h.helpTextValue
}

func TestHelperInterface(t *testing.T) {
	t.Run("help provider basic functionality", func(t *testing.T) {
		provider := &testHelpProvider{
			nameValue:     "test-command",
			helpTextValue: "This is help text for testing purposes",
		}

		// Test the Name method
		assert.Equal(t, "test-command", provider.Name())
		assert.NotEqual(t, "wrong-command", provider.Name())

		// Test the HelpText method
		assert.Equal(t, "This is help text for testing purposes", provider.HelpText())
		assert.NotEqual(t, "Wrong help text", provider.HelpText())
		assert.Contains(t, provider.HelpText(), "testing purposes")
	})

	t.Run("multiple help providers", func(t *testing.T) {
		providers := []HelpProvider{
			&testHelpProvider{
				nameValue:     "command1",
				helpTextValue: "Help for command 1",
			},
			&testHelpProvider{
				nameValue:     "command2",
				helpTextValue: "Help for command 2",
			},
		}

		// Verify each provider has the correct name and help text
		for i, provider := range providers {
			expectedName := fmt.Sprintf("command%d", i+1)
			expectedHelp := fmt.Sprintf("Help for command %d", i+1)
			
			assert.Equal(t, expectedName, provider.Name())
			assert.Equal(t, expectedHelp, provider.HelpText())
		}
	})

	t.Run("empty help text", func(t *testing.T) {
		provider := &testHelpProvider{
			nameValue:     "empty-help",
			helpTextValue: "",
		}

		assert.Equal(t, "empty-help", provider.Name())
		assert.Equal(t, "", provider.HelpText())
	})
}