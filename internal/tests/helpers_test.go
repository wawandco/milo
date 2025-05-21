package tests

import (
	"errors"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

func TestAssertNoError(t *testing.T) {
	// This should pass
	AssertNoError(t, nil)

	// We can't test the failure case directly since it would fail this test
}

func TestAssertError(t *testing.T) {
	// This should pass
	AssertError(t, errors.New("some error"))

	// We can't test the failure case directly since it would fail this test
}

func TestAssertErrorContains(t *testing.T) {
	// This should pass
	AssertErrorContains(t, errors.New("some detailed error message"), "detailed")

	// We can't test the failure cases directly since they would fail this test
}

func TestAssertEqual(t *testing.T) {
	// These should pass
	AssertEqual(t, 1, 1)
	AssertEqual(t, "test", "test")
	AssertEqual(t, []string{"a", "b"}, []string{"a", "b"})
	AssertEqual(t, map[string]int{"a": 1}, map[string]int{"a": 1})

	// We can't test the failure case directly since it would fail this test
}

func TestAssertNotEqual(t *testing.T) {
	// These should pass
	AssertNotEqual(t, 1, 2)
	AssertNotEqual(t, "test", "other")
	AssertNotEqual(t, []string{"a", "b"}, []string{"b", "a"})
	AssertNotEqual(t, map[string]int{"a": 1}, map[string]int{"b": 1})

	// We can't test the failure case directly since it would fail this test
}

func TestAssertTrue(t *testing.T) {
	// This should pass
	AssertTrue(t, true)
	AssertTrue(t, 1 == 1)
	AssertTrue(t, "a" == "a")

	// We can't test the failure case directly since it would fail this test
}

func TestAssertFalse(t *testing.T) {
	// This should pass
	AssertFalse(t, false)
	AssertFalse(t, 1 != 1)
	AssertFalse(t, "a" != "a")

	// We can't test the failure case directly since it would fail this test
}

func TestAssertContains(t *testing.T) {
	// This should pass
	AssertContains(t, "this is a test", "test")
	AssertContains(t, "abcdef", "cd")

	// We can't test the failure case directly since it would fail this test
}

func TestAssertFaults(t *testing.T) {
	// Define some test faults
	expected := []reviewers.Fault{
		{
			Reviewer: "test-reviewer",
			Line:     10,
			Col:      5,
			Rule: reviewers.Rule{
				Code:        "TEST-001",
				Description: "Test rule",
			},
			Path: "test.html",
		},
		{
			Reviewer: "another-reviewer",
			Line:     20,
			Col:      15,
			Rule: reviewers.Rule{
				Code:        "TEST-002",
				Description: "Another test rule",
			},
			Path: "test.html",
		},
	}

	// This should pass
	AssertFaults(t, expected, expected)

	// We can't test the failure cases directly since they would fail this test
}

func TestFormatMessage(t *testing.T) {
	// Test with no arguments
	result := formatMessage()
	if result != "" {
		t.Errorf("Expected empty string, got %q", result)
	}

	// Test with one string argument
	result = formatMessage("test message")
	if result != "test message: " {
		t.Errorf("Expected 'test message: ', got %q", result)
	}

	// Test with format string and arguments
	result = formatMessage("value: %d", 42)
	if result != "value: 42: " {
		t.Errorf("Expected 'value: 42: ', got %q", result)
	}

	// Test with non-string first argument
	result = formatMessage(123)
	if result != "123: " {
		t.Errorf("Expected '123: ', got %q", result)
	}
}