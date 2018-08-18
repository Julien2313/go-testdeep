// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdMapEach struct {
	BaseOKNil
	expected reflect.Value
}

var _ TestDeep = &tdMapEach{}

// MapEach operator has to be applied on maps. It compares each value
// of data map against expected value. During a match, all values have
// to match to succeed.
func MapEach(expectedValue interface{}) TestDeep {
	return &tdMapEach{
		BaseOKNil: NewBaseOKNil(3),
		expected:  reflect.ValueOf(expectedValue),
	}
}

func (m *tdMapEach) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if !got.IsValid() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "nil value",
			Got:      types.RawString("nil"),
			Expected: types.RawString("Map OR *Map"),
		})
	}

	switch got.Kind() {
	case reflect.Ptr:
		gotElem := got.Elem()
		if !gotElem.IsValid() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "nil pointer",
				Got:      types.RawString("nil " + got.Type().String()),
				Expected: types.RawString("Map OR *Map"),
			})
		}

		if gotElem.Kind() != reflect.Map {
			break
		}
		got = gotElem
		fallthrough

	case reflect.Map:
		var err *ctxerr.Error
		for _, key := range got.MapKeys() {
			err = deepValueEqual(ctx.AddMapKey(key), got.MapIndex(key), m.expected)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "bad type",
		Got:      types.RawString(got.Type().String()),
		Expected: types.RawString("Map OR *Map"),
	})
}

func (m *tdMapEach) String() string {
	const prefix = "MapEach("

	content := util.ToString(m.expected)
	if strings.Contains(content, "\n") {
		return prefix + util.IndentString(content, "        ") + ")"
	}
	return prefix + content + ")"
}
