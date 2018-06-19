// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func formatError(t TestingT, isFatal bool, err *Error, args ...interface{}) {
	t.Helper()

	const failedTest = "Failed test"

	var buf *bytes.Buffer
	switch len(args) {
	case 0:
		buf = bytes.NewBufferString(failedTest + "\n")
	case 1:
		buf = bytes.NewBufferString(failedTest + " '")
		fmt.Fprint(buf, args[0])
		buf.WriteString("'\n")
	default:
		buf = bytes.NewBufferString(failedTest + " '")
		if str, ok := args[0].(string); ok && strings.ContainsRune(str, '%') {
			fmt.Fprintf(buf, str, args[1:]...)
		} else {
			// create a new slice to fool govet and avoid "call has possible
			// formatting directive" errors
			fmt.Fprint(buf, args[:]...)
		}
		buf.WriteString("'\n")
	}

	err.Append(buf, "")

	if isFatal {
		t.Fatal(buf.String())
	} else {
		t.Error(buf.String())
	}
}

func cmpDeeply(ctx Context, t TestingT, got, expected interface{},
	args ...interface{}) bool {
	err := deepValueEqualFinal(ctx,
		reflect.ValueOf(got), reflect.ValueOf(expected))
	if err == nil {
		return true
	}
	t.Helper()
	formatError(t, ctx.failureIsFatal, err, args...)
	return false
}

// CmpDeeply returns true if "got" matches "expected". "expected" can
// be the same type as "got" is, or contains some TestDeep
// operators. If "got" does not match "expected", it returns false and
// the reason of failure is logged with the help of "t" Error()
// method.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpDeeply(t TestingT, got, expected interface{},
	args ...interface{}) bool {
	t.Helper()
	return cmpDeeply(NewContext(), t, got, expected, args...)
}
