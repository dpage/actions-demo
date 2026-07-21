// Command actions-demo is a tiny CLI wrapped around the calc and textutil
// packages; it exists purely to give the GitHub Actions demo something real to
// build, test and occasionally break.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/dpage/actions-demo/internal/calc"
	"github.com/dpage/actions-demo/internal/textutil"
)

const usage = `actions-demo - a small toy CLI for demonstrating CI

Usage:
  actions-demo add <a> <b>          add two integers
  actions-demo sub <a> <b>          subtract b from a
  actions-demo mul <a> <b>          multiply two integers
  actions-demo div <a> <b>          divide a by b
  actions-demo sum <n>...           total a list of integers
  actions-demo reverse <text>...    reverse the given text
  actions-demo palindrome <text>... report whether the text is a palindrome
  actions-demo words <text>...      count the words in the given text
  actions-demo title <text>...      title-case the given text
`

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// run executes a single command and writes its result to out, so that the CLI
// behaviour can be exercised from tests without spawning a process.
func run(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(out, usage)
		return nil
	}

	command, rest := args[0], args[1:]

	switch command {
	case "add", "sub", "mul", "div":
		a, b, err := twoInts(rest)
		if err != nil {
			return err
		}
		return arithmetic(command, a, b, out)

	case "sum":
		values, err := manyInts(rest)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, calc.Sum(values))
		return nil

	case "reverse":
		fmt.Fprintln(out, textutil.Reverse(strings.Join(rest, " ")))
		return nil

	case "palindrome":
		fmt.Fprintln(out, textutil.IsPalindrome(strings.Join(rest, " ")))
		return nil

	case "words":
		fmt.Fprintln(out, textutil.WordCount(strings.Join(rest, " ")))
		return nil

	case "title":
		fmt.Fprintln(out, textutil.Title(strings.Join(rest, " ")))
		return nil

	case "help", "-h", "--help":
		fmt.Fprint(out, usage)
		return nil

	default:
		return fmt.Errorf("unknown command %q; try 'actions-demo help'", command)
	}
}

func arithmetic(command string, a, b int, out io.Writer) error {
	switch command {
	case "add":
		fmt.Fprintln(out, calc.Add(a, b))
	case "sub":
		fmt.Fprintln(out, calc.Subtract(a, b))
	case "mul":
		fmt.Fprintln(out, calc.Multiply(a, b))
	case "div":
		result, err := calc.Divide(a, b)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, result)
	}
	return nil
}

func twoInts(args []string) (int, int, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	values, err := manyInts(args)
	if err != nil {
		return 0, 0, err
	}
	return values[0], values[1], nil
}

func manyInts(args []string) ([]int, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("expected at least one number")
	}
	values := make([]int, 0, len(args))
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("%q is not an integer", arg)
		}
		values = append(values, value)
	}
	return values, nil
}
