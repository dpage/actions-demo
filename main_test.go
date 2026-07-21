package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCommands(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"add", []string{"add", "2", "3"}, "5\n"},
		{"sub", []string{"sub", "9", "4"}, "5\n"},
		{"mul", []string{"mul", "6", "7"}, "42\n"},
		{"div", []string{"div", "84", "2"}, "42\n"},
		{"sum", []string{"sum", "1", "2", "3", "4"}, "10\n"},
		{"reverse", []string{"reverse", "hello"}, "olleh\n"},
		{"palindrome true", []string{"palindrome", "race", "car"}, "true\n"},
		{"palindrome false", []string{"palindrome", "hello"}, "false\n"},
		{"words", []string{"words", "the", "quick", "brown", "fox"}, "4\n"},
		{"title", []string{"title", "hello", "wide", "world"}, "Hello Wide World\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := run(tc.args, &out); err != nil {
				t.Fatalf("run(%v) returned unexpected error: %v", tc.args, err)
			}
			if got := out.String(); got != tc.want {
				t.Errorf("run(%v) wrote %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}

func TestRunPrintsUsage(t *testing.T) {
	for _, args := range [][]string{nil, {"help"}, {"--help"}} {
		var out bytes.Buffer
		if err := run(args, &out); err != nil {
			t.Fatalf("run(%v) returned unexpected error: %v", args, err)
		}
		if !strings.HasPrefix(out.String(), "actions-demo -") {
			t.Errorf("run(%v) did not print the usage text, got %q", args, out.String())
		}
	}
}

func TestRunErrors(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"unknown command", []string{"frobnicate"}, "unknown command"},
		{"too few arguments", []string{"add", "1"}, "expected 2 arguments"},
		{"not an integer", []string{"add", "1", "banana"}, "not an integer"},
		{"no numbers to sum", []string{"sum"}, "at least one number"},
		{"division by zero", []string{"div", "1", "0"}, "division by zero"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(tc.args, &out)
			if err == nil {
				t.Fatalf("run(%v) returned nil, want an error", tc.args)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("run(%v) returned error %q, want it to contain %q", tc.args, err, tc.want)
			}
		})
	}
}
