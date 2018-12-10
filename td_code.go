// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdCode struct {
	Base
	function reflect.Value
	argType  reflect.Type
}

var _ TestDeep = &tdCode{}

// Code operator allows to check data using a custom function. So
// "fn" is a function that must take one parameter whose type must be
// the same as the type of the compared value.
//
// "fn" can return a single bool kind value, telling that yes or no
// the custom test is successful:
//   Code(func (date time.Time) bool {
//       return date.Year() == 2018
//     })
//
// or two values (bool, string) kinds. The bool value has the same
// meaning as above, and the string value is used to describe the
// test when it fails:
//   Code(func (date time.Time) (bool, string) {
//       if date.Year() == 2018 {
//         return true, ""
//       }
//       return false, "year must be 2018"
//     })
//
// or a single error value. If the returned error is nil, the test
// succeeded, else the error contains the reason of failure:
//   Code(func (b json.RawMessage) error {
//       var c map[string]int
//       err := json.Unmarshal(b, &c)
//       if err != nil {
//         return err
//       }
//       if c["test"] != 42 {
//         return fmt.Errorf(`key "test" does not match 42`)
//       }
//       return nil
//     })
//
// This operator allows to handle any specific comparison not handled
// by standard operators.
//
// It is not recommended to call CmpDeeply (or any other Cmp*
// functions or *T methods) inside the body of "fn", because of
// confusion produced by output in case of failure. When the data
// needs to be transformed before being compared again, Smuggle
// operator should be used instead.
//
// TypeBehind method returns the reflect.Type of only parameter of "fn".
func Code(fn interface{}) TestDeep {
	vfn := reflect.ValueOf(fn)

	if vfn.Kind() != reflect.Func {
		panic("usage: Code(FUNC)")
	}

	fnType := vfn.Type()
	if fnType.NumIn() != 1 {
		panic("Code(FUNC): FUNC must take only one argument")
	}

	switch fnType.NumOut() {
	case 2:
		if fnType.Out(1).Kind() != reflect.String {
			break
		}
		fallthrough

	case 1:
		// (*bool*) or (*bool*, string)
		if fnType.Out(0).Kind() == reflect.Bool ||
			// (*error*)
			(fnType.NumOut() == 1 && fnType.Out(0) == errorInterface) {
			return &tdCode{
				Base:     NewBase(3),
				function: vfn,
				argType:  fnType.In(0),
			}
		}
	}

	panic("Code(FUNC): FUNC must return bool or (bool, string) or error")
}

func (c *tdCode) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if !got.Type().AssignableTo(c.argType) {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "incompatible parameter type",
			Got:      types.RawString(got.Type().String()),
			Expected: types.RawString(c.argType.String()),
		})
	}

	// Refuse to override unexported fields access in this case. It is a
	// choice, as we think it is better to use Code() on surrounding
	// struct instead.
	if !got.CanInterface() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot compare unexported field",
			Summary: types.RawString("use Code() on surrounding struct instead"),
		})
	}

	ret := c.function.Call([]reflect.Value{got})
	if ret[0].Kind() == reflect.Bool {
		if ret[0].Bool() {
			return nil
		}
	} else if ret[0].IsNil() { // reflect.Interface
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	summary := tdCodeResult{
		Value: got,
	}

	if len(ret) > 1 { // (bool, string)
		summary.Reason = ret[1].String()
	} else if ret[0].Kind() == reflect.Interface { // (error)
		summary.Reason = ret[0].Interface().(error).Error()
	}
	// else (bool) so no reason to report

	return ctx.CollectError(&ctxerr.Error{
		Message: "ran code with %% as argument",
		Summary: summary,
	})
}

func (c *tdCode) String() string {
	return "Code(" + c.function.Type().String() + ")"
}

func (c *tdCode) TypeBehind() reflect.Type {
	return c.argType
}

type tdCodeResult struct {
	types.TestDeepStamp
	Value  reflect.Value
	Reason string
}

var _ types.TestDeepStringer = tdCodeResult{}

func (r tdCodeResult) String() string {
	if r.Reason == "" {
		return fmt.Sprintf("  value: %s\nit failed but didn't say why",
			util.IndentString(util.ToString(r.Value), "         "))
	}
	return fmt.Sprintf("        value: %s\nit failed coz: %s",
		util.IndentString(util.ToString(r.Value), "               "), r.Reason)
}
