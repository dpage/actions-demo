# actions-demo

[![CI](https://github.com/dpage/actions-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/dpage/actions-demo/actions/workflows/ci.yml)

A deliberately small Go CLI whose only real job is to give a GitHub Actions
demo something to build, test, and occasionally fall over on. There are two
tiny packages under `internal/`, a thin command-line wrapper in `main.go`, and
a CI workflow that lints, tests across two Go versions, and then builds and
smoke-tests the resulting binary.

## Building and running

```console
$ go build -o actions-demo .
$ ./actions-demo add 2 3
5
$ ./actions-demo palindrome "A man, a plan, a canal: Panama"
true
$ ./actions-demo title "hello wide world"
Hello Wide World
```

Run `./actions-demo help` for the full list of commands.

## Running the tests

```console
$ go test ./...
$ go test -race -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
```

## The CI workflow

`.github/workflows/ci.yml` runs on every push and pull request against `main`,
and can also be triggered by hand from the Actions tab. It has three jobs:

- **Lint**, which fails if anything is not `gofmt`'d, and then runs `go vet`.
- **Test**, a matrix across Go 1.25 and the current stable release, running the
  tests with the race detector and printing a coverage summary.
- **Build**, which depends on the other two, compiles the binary, and smoke
  tests it by checking that `add 2 3` really does return 5.

## Breaking it on purpose

The whole point of this repository is that it can be broken and unbroken in a
single edit, so here are a few reliable ways to do that, in rough order of how
quickly the failure shows up.

| What to change | Which job fails | What you will see |
|---|---|---|
| In `internal/calc/calc.go`, make `Add` return `a - b` | Test, then Build | Four subtests of `TestAdd` fail, along with the `add` case in `TestRunCommands` |
| In `internal/calc/calc.go`, drop the `b == 0` guard in `Divide` | Test | `TestDivideByZero` fails, and the CLI panics with an integer divide by zero |
| In `internal/textutil/textutil.go`, change `Reverse` to iterate over `[]byte` rather than `[]rune` | Test | Only the multi-byte case fails, which is a nice illustration of a test earning its keep |
| In `internal/textutil/textutil.go`, remove the `unicode.IsDigit` check in `Normalise` | Test | `TestNormalise` fails on the digits case whilst everything else passes |
| Add a stray `if x := 1; x == 1 {` with an unused variable, or an unreachable `return` | Lint | `go vet` complains before the tests ever run |
| Indent something with spaces instead of a tab | Lint | The `gofmt -l` step lists the offending file |
| In `internal/calc/calc.go`, make `Max` start from `0` rather than `values[0]` | Test | Only the all-negative case fails |

The last two are particularly useful if you want to show the difference between
a failure that a formatter can fix mechanically and one that needs somebody to
think about the logic.

To unbreak, revert the edit; `git checkout -- .` or `git revert` both work, and
pushing the fix will turn the badge green again within a minute or so.
