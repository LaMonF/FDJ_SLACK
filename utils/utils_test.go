package utils

import (
	"testing"
)

func TestUtils_IsWhitespace(t *testing.T) {
	// One white space Expect true
	result := IsWhitespace(" ")
	if result != true {
		t.Fail()
	}

	// multiple white space Expect true
	result = IsWhitespace("      ")
	if result != true {
		t.Fail()
	}

	// No spaces expect false
	result = IsWhitespace("a")
	if result != false {
		t.Fail()
	}

	// spaces before and ending with letter expect false
	result = IsWhitespace("    a")
	if result != false {
		t.Fail()
	}

	// space after starting with a letter expect false
	result = IsWhitespace("q       ")
	if result != false {
		t.Fail()
	}

	// multiple letters with space in middle expect false
	result = IsWhitespace("aqq  qqqq")
	if result != false {
		t.Fail()
	}
}

