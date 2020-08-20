package errand

import (
	"fmt"
	"strings"
)

// Errors is responsible for listing multiple errors.
type Errors interface {
	// Errors returns the list of wrapped errors.
	Errors() []error
}

// errand offers the error slice.
type errand []error

// Error implements the error interface.
func (e errand) Error() string {
	// Technically we should never have an errand with length 0 or 1
	s := make([]string, len(e))

	for i, err := range e {
		s[i] = err.Error()
	}

	return fmt.Sprintf("%d errors: %s", len(s), strings.Join(s, ", "))
}

// Errors returns the slice of errors.
func (e errand) Errors() []error {
	return e
}

// Append errs to err. Any nil error is ignored. If only one error is left, that error is exactly returned. If none are left, err is returned (so it keeps the type). If any of the err or errs is an errand, only their entries are taken and the provided order is kept.
func Append(err error, errs ...error) error {
	errs = append([]error{err}, errs...)

	var r errand

	for _, e := range errs {
		if e != nil {
			switch t := e.(type) {
			// if e is an errand, take only its items
			case errand:
				r = append(r, t...)
			default:
				r = append(r, e)
			}
		}
	}

	if len(r) == 0 {
		// to keep the type of err, still nil
		return err
	}

	if len(r) == 1 {
		return r[0]
	}

	return r
}
