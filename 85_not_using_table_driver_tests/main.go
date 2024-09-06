package main

import (
	"strings"
	"testing"
)

func main() {

}

// go test -run=TestFoo/subtest_1 -v      run only subtest 1
func TestFoo(t *testing.T) {
	t.Run("subtest 1", func(t *testing.T) {
		if false {
			t.Error()
		}
	})
	t.Run("subtest 2", func(t *testing.T) {
		if 2 != 2 {
			t.Error()
		}
	})
}

func TestRemoveNewLineSuffix(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		`empty`: {
			input:    "",
			expected: "",
		},
		`ending with \r\n`: {
			input:    "a\r\n",
			expected: "a",
		},
		`ending with \n`: {
			input:    "a\n",
			expected: "a",
		},
		`ending with multiple \n`: {
			input:    "a\n\n\n",
			expected: "a",
		},
		`ending without newline`: {
			input:    "a",
			expected: "a",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := removeNewLineSuffixes(tt.input)
			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}

func removeNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "\r\n") {
		return removeNewLineSuffixes(s[:len(s)-2])
	}
	if strings.HasSuffix(s, "\n") {
		return removeNewLineSuffixes(s[:len(s)-1])
	}
	return s
}
