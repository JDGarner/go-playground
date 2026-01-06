package intervals

import "testing"

func TestAddBoldTag(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		words  []string
		expect string
	}{
		// Basic cases
		{
			name:   "non-overlapping substrings",
			s:      "abcxyz123",
			words:  []string{"abc", "123"},
			expect: "<b>abc</b>xyz<b>123</b>",
		},
		{
			name:   "overlapping substrings",
			s:      "aaabbb",
			words:  []string{"aa", "b"},
			expect: "<b>aaabbb</b>",
		},

		// Edge cases - empty inputs
		{
			name:   "empty string",
			s:      "",
			words:  []string{"abc"},
			expect: "",
		},
		{
			name:   "empty words array",
			s:      "abcdef",
			words:  []string{},
			expect: "abcdef",
		},
		{
			name:   "both empty",
			s:      "",
			words:  []string{},
			expect: "",
		},

		// No matches
		{
			name:   "no word found in string",
			s:      "abcdef",
			words:  []string{"xyz", "123"},
			expect: "abcdef",
		},

		// Single character cases
		{
			name:   "single character match",
			s:      "a",
			words:  []string{"a"},
			expect: "<b>a</b>",
		},
		{
			name:   "single character no match",
			s:      "a",
			words:  []string{"b"},
			expect: "a",
		},

		// Consecutive substrings (should merge)
		{
			name:   "consecutive substrings",
			s:      "abcdef",
			words:  []string{"abc", "def"},
			expect: "<b>abcdef</b>",
		},
		{
			name:   "consecutive with gap",
			s:      "abcxdef",
			words:  []string{"abc", "def"},
			expect: "<b>abc</b>x<b>def</b>",
		},

		// Overlapping cases
		{
			name:   "partially overlapping",
			s:      "abcdef",
			words:  []string{"abc", "cde"},
			expect: "<b>abcde</b>f",
		},
		{
			name:   "fully contained substring",
			s:      "abcdef",
			words:  []string{"abcdef", "cd"},
			expect: "<b>abcdef</b>",
		},
		{
			name:   "multiple overlaps",
			s:      "aaaa",
			words:  []string{"a", "aa", "aaa"},
			expect: "<b>aaaa</b>",
		},

		// Multiple occurrences
		{
			name:   "word appears multiple times",
			s:      "abcabc",
			words:  []string{"abc"},
			expect: "<b>abcabc</b>",
		},
		{
			name:   "word appears multiple times with gap",
			s:      "abcxabc",
			words:  []string{"abc"},
			expect: "<b>abc</b>x<b>abc</b>",
		},
		{
			name:   "multiple words multiple occurrences",
			s:      "ababab",
			words:  []string{"a", "b"},
			expect: "<b>ababab</b>",
		},

		// Substring at boundaries
		{
			name:   "word at start",
			s:      "abcdef",
			words:  []string{"abc"},
			expect: "<b>abc</b>def",
		},
		{
			name:   "word at end",
			s:      "abcdef",
			words:  []string{"def"},
			expect: "abc<b>def</b>",
		},
		{
			name:   "entire string is one word",
			s:      "abcdef",
			words:  []string{"abcdef"},
			expect: "<b>abcdef</b>",
		},

		// Complex overlapping scenarios
		{
			name:   "three overlapping words",
			s:      "abcdefgh",
			words:  []string{"abc", "cde", "efg"},
			expect: "<b>abcdefg</b>h",
		},
		{
			name:   "chain of consecutive",
			s:      "abcdefghij",
			words:  []string{"ab", "cd", "ef", "gh"},
			expect: "<b>abcdefgh</b>ij",
		},

		// Duplicate words in array
		{
			name:   "duplicate words in array",
			s:      "abcxxyz",
			words:  []string{"abc", "abc", "xyz"},
			expect: "<b>abc</b>x<b>xyz</b>",
		},

		// Words longer than string
		{
			name:   "word longer than string",
			s:      "abc",
			words:  []string{"abcdef"},
			expect: "abc",
		},

		// Special patterns
		{
			name:   "repeating pattern",
			s:      "ababababab",
			words:  []string{"ab"},
			expect: "<b>ababababab</b>",
		},
		{
			name:   "interleaved matches",
			s:      "axbxcxdx",
			words:  []string{"a", "b", "c", "d"},
			expect: "<b>a</b>x<b>b</b>x<b>c</b>x<b>d</b>x",
		},

		// Adjacent but not overlapping
		{
			name:   "adjacent intervals exactly",
			s:      "abcd",
			words:  []string{"ab", "cd"},
			expect: "<b>abcd</b>",
		},

		// Case sensitivity (assuming case-sensitive matching)
		{
			name:   "case sensitive - no match",
			s:      "AbC",
			words:  []string{"abc"},
			expect: "AbC",
		},
		{
			name:   "case sensitive - exact match",
			s:      "AbC",
			words:  []string{"AbC"},
			expect: "<b>AbC</b>",
		},

		// Numbers and special characters
		{
			name:   "with numbers",
			s:      "abc123def456",
			words:  []string{"123", "456"},
			expect: "abc<b>123</b>def<b>456</b>",
		},
		{
			name:   "with special characters",
			s:      "hello-world!",
			words:  []string{"hello", "world"},
			expect: "<b>hello</b>-<b>world</b>!",
		},

		// Complex merge scenarios
		{
			name:   "multiple merges needed",
			s:      "aaaaaa",
			words:  []string{"aa", "aaa"},
			expect: "<b>aaaaaa</b>",
		},
		{
			name:   "gaps between some intervals",
			s:      "abcxdefxyzhij",
			words:  []string{"abc", "def", "hij"},
			expect: "<b>abc</b>x<b>def</b>xyz<b>hij</b>",
		},
		{
			name:   "single char gaps",
			s:      "axbxc",
			words:  []string{"a", "b", "c"},
			expect: "<b>a</b>x<b>b</b>x<b>c</b>",
		},

		// Stress patterns
		{
			name:   "all single characters",
			s:      "abcdefgh",
			words:  []string{"a", "b", "c", "d", "e", "f", "g", "h"},
			expect: "<b>abcdefgh</b>",
		},
		{
			name:   "alternating pattern",
			s:      "ababab",
			words:  []string{"ab"},
			expect: "<b>ababab</b>",
		},
		{
			name:   "unicode characters",
			s:      "你好xab",
			words:  []string{"好", "ab"},
			expect: "你<b>好</b>x<b>ab</b>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddBoldTag(tt.s, tt.words)
			if result != tt.expect {
				t.Errorf("AddBoldTag(%q, %v) = %q; want %q", tt.s, tt.words, result, tt.expect)
			}
		})
	}
}
