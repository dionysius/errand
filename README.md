# errand

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dionysius/errand)](https://github.com/dionysius/errand/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/dionysius/errand)](https://goreportcard.com/report/github.com/dionysius/errand)
[![GitHub](https://img.shields.io/github/license/dionysius/errand?color=lightgrey)](https://github.com/dionysius/errand/blob/master/COPYING)

errand is your little assistant for error slices or also called multierrors. It helps you append multiple errors together into a single error as minimally as possible. The usage is as native as possible to the original error interface.

errand is the right library for you, if you just want a slice of actual errors from a bunch of calls. It gives you the following features:

- you can directly use errand - no need to initialize the error slice
- appending nil doesn't add another item to the slice - no need to check for nil before appending
- if the error you append to and append from is nil, the error is still nil - no need for a special method to get ErrOrNil() for your returns
- it doesn't force you that error is an error slice - returns that exact error if only one error was appended
- if any error is already an errand, only the entries are taken - exactly one level of error slice is returned, you are flexible in how you organize your code and where you return parts of your error slice
- no external dependencies - it doesn't add weight

errand is **not a replacement for wrapping**. I encourage you to use [golangs wrapping functionality](https://blog.golang.org/go1.13-errors)! As of now, you can't find out the types of these errors.

## Usage

errand offers only one package exported function: `Append` (besides the error and string interface type implementations `Error` and `String`).

To download errand:

`go get github.com/dionysius/errand`

To import errand:

`import "github.com/dionysius/errand`

And use errand like you would with go's append:

```go
package main

import (
  "fmt"

  errand "github.com/dionysius/errand"
)

func main() {
  errs := fmt.Errorf("a function with possibly an error")

  errs = errand.Append(errs, fmt.Errorf("directly add a possible error"))

  // multiple errands are merged together, we get an errand from the separate function
  errs = errand.Append(errs, separate())

  // any nil errors are ignored
  errs = errand.Append(errs, nil)

  errs = errand.Append(errs, willBeNil())

  // Output the error
  fmt.Printf("%s\n", errs)
}

func separate() error {
  errs := errand.Append(nil, fmt.Errorf("you can directly use errand without initialisation of the first error"))

  // If there is only one error, like the command above, errs will be exactly that error, not an errand error slice.

  return errand.Append(errs, fmt.Errorf("and force that this function returns an errand since two errors were added"))
}

func willBeNil() error {
  var err error

  // If everything is nil, Append returns nil
  return errand.Append(err, nil)
}
```

The output of the above example is (test it yourself with `go run internal/readme.go`):

```txt
4 errors: a function with possibly an error, directly add a possible error, you can directly use errand without initialisation of the first error, and force that this function returns an errand since two errors were added
```

## Todos

- Implement the interfaces for `errors.Is` and `errors.As`. PRs welcome!
- Not sure yet if I even want to: But we could offer an interface `Errander` to offer implementations `Append(error, []error...)` so they can benefit from the merging of error slices into errand.

## Motivation

Most multierror packages are quite heavy. While they offer many features, It's often more than needed and that complicates their usage more than it should be. Don't get me wrong, they're definitely helpful and useful if you need those features!

errand is inspired by [multierror from felixge](https://github.com/felixge/multierror). But since there are no licence informations, looks abandoned and misses a tiny little feature to append two error slices together, I've decided to create errand.
