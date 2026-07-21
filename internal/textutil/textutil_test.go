package textutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"simple word", "hello", "olleh"},
		{"multi-byte runes", "héllo", "olléh"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.in); got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestNormalise(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"already clean", "abc", "abc"},
		{"mixed case", "AbC", "abc"},
		{"punctuation stripped", "a, b. c!", "abc"},
		{"digits kept", "a1 b2", "a1b2"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Normalise(tc.in); got != tc.want {
				t.Errorf("Normalise(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"empty string", "", true},
		{"single character", "x", true},
		{"simple palindrome", "racecar", true},
		{"cased palindrome", "RaceCar", true},
		{"punctuated palindrome", "A man, a plan, a canal: Panama", true},
		{"not a palindrome", "hello", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.in); got != tc.want {
				t.Errorf("IsPalindrome(%q) = %t, want %t", tc.in, got, tc.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{"empty string", "", 0},
		{"only whitespace", "   \t\n ", 0},
		{"single word", "hello", 1},
		{"several words", "the quick brown fox", 4},
		{"irregular spacing", "  one   two  ", 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordCount(tc.in); got != tc.want {
				t.Errorf("WordCount(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
		{"several words", "hello wide world", "Hello Wide World"},
		{"existing capitals preserved", "hello wIde world", "Hello WIde World"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Title(tc.in); got != tc.want {
				t.Errorf("Title(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
