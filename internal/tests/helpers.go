package tests

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/wawandco/milo/reviewers"
)

// AssertNoError asserts that the provided error is nil. 
// Fails the test with a descriptive message if err is not nil.
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: unexpected error: %v", filepath.Base(file), line, err)
	}
}

// AssertError asserts that the provided error is not nil.
// Fails the test if err is nil.
func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: expected error but got nil", filepath.Base(file), line)
	}
}

// AssertErrorContains asserts that the provided error is not nil and contains the specified substring.
// Fails the test if err is nil or if it doesn't contain the expected substring.
func AssertErrorContains(t *testing.T, err error, substring string) {
	t.Helper()
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: expected error containing %q but got nil", filepath.Base(file), line, substring)
	}
	
	if !strings.Contains(err.Error(), substring) {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: expected error containing %q but got %q", filepath.Base(file), line, substring, err.Error())
	}
}

// AssertEqual asserts that got equals want, using reflect.DeepEqual for comparison.
// Fails the test with a descriptive message if the values are not equal.
func AssertEqual(t *testing.T, got, want interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		msg := formatMessage(msgAndArgs...)
		t.Fatalf("%s:%d: %sexpected %v, got %v", filepath.Base(file), line, msg, want, got)
	}
}

// AssertNotEqual asserts that got does not equal want, using reflect.DeepEqual for comparison.
// Fails the test with a descriptive message if the values are equal.
func AssertNotEqual(t *testing.T, got, want interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		msg := formatMessage(msgAndArgs...)
		t.Fatalf("%s:%d: %sexpr: %v should not equal %v", filepath.Base(file), line, msg, got, want)
	}
}

// AssertTrue asserts that the provided value is true.
// Fails the test with a descriptive message if the value is not true.
func AssertTrue(t *testing.T, value bool, msgAndArgs ...interface{}) {
	t.Helper()
	if !value {
		_, file, line, _ := runtime.Caller(1)
		msg := formatMessage(msgAndArgs...)
		t.Fatalf("%s:%d: %sassertion failed: expected true but got false", filepath.Base(file), line, msg)
	}
}

// AssertFalse asserts that the provided value is false.
// Fails the test with a descriptive message if the value is not false.
func AssertFalse(t *testing.T, value bool, msgAndArgs ...interface{}) {
	t.Helper()
	if value {
		_, file, line, _ := runtime.Caller(1)
		msg := formatMessage(msgAndArgs...)
		t.Fatalf("%s:%d: %sassertion failed: expected false but got true", filepath.Base(file), line, msg)
	}
}

// AssertContains asserts that substring is a substring of str.
// Fails the test with a descriptive message if substring is not found in str.
func AssertContains(t *testing.T, str, substring string, msgAndArgs ...interface{}) {
	t.Helper()
	if !strings.Contains(str, substring) {
		_, file, line, _ := runtime.Caller(1)
		msg := formatMessage(msgAndArgs...)
		t.Fatalf("%s:%d: %sexpected %q to contain %q", filepath.Base(file), line, msg, str, substring)
	}
}

// AssertFaults checks if the actual faults match the expected faults.
// This is specific to the Milo reviewers package.
func AssertFaults(t *testing.T, got, want []reviewers.Fault) {
	t.Helper()
	if len(got) != len(want) {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: expected %d faults, got %d", filepath.Base(file), line, len(want), len(got))
	}
	
	for i, wantFault := range want {
		if i >= len(got) {
			_, file, line, _ := runtime.Caller(1)
			t.Fatalf("%s:%d: missing expected fault at index %d", filepath.Base(file), line, i)
		}
		
		gotFault := got[i]
		if gotFault.Reviewer != wantFault.Reviewer {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Reviewer %s, got %s", 
				filepath.Base(file), line, i, wantFault.Reviewer, gotFault.Reviewer)
		}
		if gotFault.Line != wantFault.Line {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Line %d, got %d", 
				filepath.Base(file), line, i, wantFault.Line, gotFault.Line)
		}
		if gotFault.Col != wantFault.Col {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Col %d, got %d", 
				filepath.Base(file), line, i, wantFault.Col, gotFault.Col)
		}
		if wantFault.Rule.Code != "" && gotFault.Rule.Code != wantFault.Rule.Code {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Rule.Code %s, got %s", 
				filepath.Base(file), line, i, wantFault.Rule.Code, gotFault.Rule.Code)
		}
		if wantFault.Rule.Description != "" && gotFault.Rule.Description != wantFault.Rule.Description {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Rule.Description %s, got %s", 
				filepath.Base(file), line, i, wantFault.Rule.Description, gotFault.Rule.Description)
		}
		if wantFault.Path != "" && gotFault.Path != wantFault.Path {
			_, file, line, _ := runtime.Caller(1)
			t.Errorf("%s:%d: fault[%d]: expected Path %s, got %s", 
				filepath.Base(file), line, i, wantFault.Path, gotFault.Path)
		}
	}
}

// formatMessage formats additional message arguments for assertion failures.
// If no arguments are provided, returns an empty string.
// Otherwise, returns a formatted string with a trailing space.
func formatMessage(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	
	msg := msgAndArgs[0]
	args := msgAndArgs[1:]
	
	if msg, ok := msg.(string); ok {
		return fmt.Sprintf(msg, args...) + ": "
	}
	
	return fmt.Sprintf("%v: ", msg)
}