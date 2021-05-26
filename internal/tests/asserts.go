package tests

import (
	"github.com/ozoncp/ocp-issue-api/internal/issues"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func assertEqualSlice(t *testing.T, expected []int, actual []int) {
	if len(expected) != len(actual) {
		t.Errorf("expected size: %d, actual size: %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected: %d, actual: %d at index: %d", expected[i], actual[i], i)
		}
	}
}

func assertEqualMap(t *testing.T, expected map[string]int, actual map[string]int) {
	if len(expected) != len(actual) {
		t.Errorf("expected size: %d, actual size: %d", len(expected), len(actual))
	}

	for key, value := range expected {
		if value != actual[key] {
			t.Errorf("expected value: %d, actual value: %d for key: %s", value, actual[key], key)
		}
	}
}

func assertEqualIssue(t *testing.T, expected issues.Issue, actual issues.Issue) {
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected.String(), actual.String())
	}
}

func assertEqualIssueSlice(t *testing.T, expected []issues.Issue, actual []issues.Issue) {
	if len(expected) != len(actual) {
		t.Errorf("expected size: %d, actual size: %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected: %s, actual: %s at index: %d", expected[i].String(), actual[i].String(), i)
		}
	}
}
