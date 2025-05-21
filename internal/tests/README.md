# Milo Test Helpers

This package provides a collection of test helper functions that simplify writing tests in the Milo project.

## Usage

```go
import (
    "testing"
    "github.com/wawandco/milo/internal/tests"
)

func TestSomething(t *testing.T) {
    // Use the helpers in your tests
    tests.AssertNoError(t, err)
    tests.AssertEqual(t, got, want)
    tests.AssertFaults(t, gotFaults, expectedFaults)
}
```

## Available Helpers

### Error Assertions

- `AssertNoError(t *testing.T, err error)` - Asserts that the error is nil
- `AssertError(t *testing.T, err error)` - Asserts that the error is not nil
- `AssertErrorContains(t *testing.T, err error, substring string)` - Asserts that the error contains a substring

### Equality Assertions

- `AssertEqual(t *testing.T, got, want interface{}, msgAndArgs ...interface{})` - Asserts that got equals want
- `AssertNotEqual(t *testing.T, got, want interface{}, msgAndArgs ...interface{})` - Asserts that got does not equal want

### Boolean Assertions

- `AssertTrue(t *testing.T, value bool, msgAndArgs ...interface{})` - Asserts that value is true
- `AssertFalse(t *testing.T, value bool, msgAndArgs ...interface{})` - Asserts that value is false

### String Assertions

- `AssertContains(t *testing.T, str, substring string, msgAndArgs ...interface{})` - Asserts str contains substring

### Milo-Specific Assertions

- `AssertFaults(t *testing.T, got, want []reviewers.Fault)` - Asserts that the actual faults match the expected faults

## Benefits

These helpers:

1. Reduce boilerplate code in tests
2. Provide consistent error messages with file and line numbers
3. Simplify common assertion patterns in the Milo project
4. Make tests more readable and maintainable

## Implementation Notes

All helper functions use `t.Helper()` to ensure that error messages point to the correct line in the test file, not within the helper functions.