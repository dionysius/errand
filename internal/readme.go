// nolint: goerr113
package main

import (
	"fmt"

	"github.com/dionysius/errand"
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

	// Or we can also get the error slice using the Errors interface
	if multi, is := errs.(errand.Errors); is {
		for _, err := range multi.Errors() {
			fmt.Println("->", err.Error())
		}
	}
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
