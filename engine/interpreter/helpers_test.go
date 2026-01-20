package interpreter

import (
	"slices"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a + b", []string{"a", "+", "b"}},
		{"a-b", []string{"a", "-", "b"}},
		{"a * b", []string{"a", "*", "b"}},
		{"a/b", []string{"a", "/", "b"}},
		{"a^b", []string{"a", "^", "b"}},
		{"a!b", []string{"a", "!", "b"}},
		{"a%b", []string{"a", "%", "b"}},
		{"a=b", []string{"a", "=", "b"}},
		{"a>b", []string{"a", ">", "b"}},
		{"a<b", []string{"a", "<", "b"}},
		{"'a b' + c", []string{"'a b'", "+", "c"}},
		{"\"a b\" + c", []string{"\"a b\"", "+", "c"}},
		{" a + b ", []string{"a", "+", "b"}},
	}

	for _, test := range tests {
		result := tokenizer(test.input)
		if !slices.Equal(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
