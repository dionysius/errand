// nolint: goerr113, testpackage
package errand

import (
	"errors"
	"fmt"
	"go/scanner"
	"testing"
)

// Guarantee errand implements Errors and error.
var (
	_ Errors = new(errand)
	_ error  = new(errand)
)

// The tests below for Append make sure there is never an errand with less than 2 entries.
func Test_errand_Error_Multi(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("one")
	err2 := fmt.Errorf("two")
	err3 := Append(err, err2)

	if fmt.Sprintf("%s", err3) != "2 errors: one, two" {
		t.Fail()
	}
}

func Test_errand_Errors(t *testing.T) {
	t.Parallel()

	errs := []error{
		fmt.Errorf("one"),
		fmt.Errorf("two"),
	}

	erra := Append(nil, errs[0])
	erra = Append(erra, errs[1])

	errand, is := erra.(Errors)
	if !is {
		t.Fail()
	}

	errb := errand.Errors()

	for i := range errs {
		if errs[i] != errb[i] {
			t.Fail()
		}
	}
}

func Test_Append_Nils(t *testing.T) {
	t.Parallel()

	err := Append(nil, nil)
	if err != nil {
		t.Fail()
	}
}

func Test_Append_NilErr(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("my error")
	err2 := Append(nil, err)

	if !errors.Is(err, err2) {
		t.Fail()
	}
}

func Test_Append_NilErrs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("my error")
	err2 := Append(err, nil)

	if !errors.Is(err, err2) {
		t.Fail()
	}
}

func Test_Append_EmptyErrs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("my error")
	err2 := Append(err, []error{}...)

	if !errors.Is(err, err2) {
		t.Fail()
	}
}

func Test_Append_MultiErrs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("one")
	err2 := fmt.Errorf("two")
	err3 := Append(nil, err, err2)

	err3and, is := err3.(errand)
	if !is {
		t.Fail()
	}

	if len(err3and) != 2 {
		t.Fail()
	}
}

func Test_Append_ErrAndErrs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("one")
	err2 := fmt.Errorf("two")
	err3 := Append(err, err2)

	err3and, is := err3.(errand)
	if !is {
		t.Fail()
	}

	if len(err3and) != 2 {
		t.Fail()
	}
}

func Test_Append_ErrAndMultiErrs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("one")
	err2 := fmt.Errorf("two")
	err3 := fmt.Errorf("three")
	err4 := Append(err, err2, err3)

	err4and, is := err4.(errand)
	if !is {
		t.Fail()
	}

	if len(err4and) != 3 {
		t.Fail()
	}
}

func Test_Append_NilErrType(t *testing.T) {
	t.Parallel()

	var err *scanner.Error
	err2 := Append(err, nil)

	if !errors.Is(err, err2) {
		t.Fail()
	}

	if fmt.Sprintf("%T", err) != fmt.Sprintf("%T", err2) {
		t.Fail()
	}
}

func Test_Append_AddErrand(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("one")
	err2 := fmt.Errorf("two")
	err3 := fmt.Errorf("three")
	err4 := Append(err, err2)
	err5 := Append(err3, err4)

	err5and, is := err5.(errand)
	if !is {
		t.Fail()
	}

	if len(err5and) != 3 {
		t.Fail()
	}
}
