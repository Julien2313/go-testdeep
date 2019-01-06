// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// DO NOT EDIT!!! AUTOMATICALLY GENERATED!!!

package testdeep

import (
	"testing" // used by t.Helper() workaround below
	"time"
)

// CmpAll is a shortcut for:
//
//   CmpDeeply(t, got, All(expectedValues...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#All for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpAll(t TestingT, got interface{}, expectedValues []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, All(expectedValues...), args...)
}

// CmpAny is a shortcut for:
//
//   CmpDeeply(t, got, Any(expectedValues...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Any for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpAny(t TestingT, got interface{}, expectedValues []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Any(expectedValues...), args...)
}

// CmpArray is a shortcut for:
//
//   CmpDeeply(t, got, Array(model, expectedEntries), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Array for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpArray(t TestingT, got interface{}, model interface{}, expectedEntries ArrayEntries, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Array(model, expectedEntries), args...)
}

// CmpArrayEach is a shortcut for:
//
//   CmpDeeply(t, got, ArrayEach(expectedValue), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#ArrayEach for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpArrayEach(t TestingT, got interface{}, expectedValue interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, ArrayEach(expectedValue), args...)
}

// CmpBag is a shortcut for:
//
//   CmpDeeply(t, got, Bag(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Bag for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpBag(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Bag(expectedItems...), args...)
}

// CmpBetween is a shortcut for:
//
//   CmpDeeply(t, got, Between(from, to, bounds), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Between for details.
//
// Between() optional parameter "bounds" is here mandatory.
// BoundsInIn value should be passed to mimic its absence in
// original Between() call.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpBetween(t TestingT, got interface{}, from interface{}, to interface{}, bounds BoundsKind, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Between(from, to, bounds), args...)
}

// CmpCap is a shortcut for:
//
//   CmpDeeply(t, got, Cap(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Cap for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpCap(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Cap(val), args...)
}

// CmpCode is a shortcut for:
//
//   CmpDeeply(t, got, Code(fn), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Code for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpCode(t TestingT, got interface{}, fn interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Code(fn), args...)
}

// CmpContains is a shortcut for:
//
//   CmpDeeply(t, got, Contains(expectedValue), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Contains for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpContains(t TestingT, got interface{}, expectedValue interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Contains(expectedValue), args...)
}

// CmpContainsKey is a shortcut for:
//
//   CmpDeeply(t, got, ContainsKey(expectedValue), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#ContainsKey for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpContainsKey(t TestingT, got interface{}, expectedValue interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, ContainsKey(expectedValue), args...)
}

// CmpEmpty is a shortcut for:
//
//   CmpDeeply(t, got, Empty(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Empty for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpEmpty(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Empty(), args...)
}

// CmpGt is a shortcut for:
//
//   CmpDeeply(t, got, Gt(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Gt for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpGt(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Gt(val), args...)
}

// CmpGte is a shortcut for:
//
//   CmpDeeply(t, got, Gte(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Gte for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpGte(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Gte(val), args...)
}

// CmpHasPrefix is a shortcut for:
//
//   CmpDeeply(t, got, HasPrefix(expected), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#HasPrefix for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpHasPrefix(t TestingT, got interface{}, expected string, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, HasPrefix(expected), args...)
}

// CmpHasSuffix is a shortcut for:
//
//   CmpDeeply(t, got, HasSuffix(expected), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#HasSuffix for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpHasSuffix(t TestingT, got interface{}, expected string, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, HasSuffix(expected), args...)
}

// CmpIsa is a shortcut for:
//
//   CmpDeeply(t, got, Isa(model), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Isa for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpIsa(t TestingT, got interface{}, model interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Isa(model), args...)
}

// CmpLen is a shortcut for:
//
//   CmpDeeply(t, got, Len(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Len for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpLen(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Len(val), args...)
}

// CmpLt is a shortcut for:
//
//   CmpDeeply(t, got, Lt(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Lt for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpLt(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Lt(val), args...)
}

// CmpLte is a shortcut for:
//
//   CmpDeeply(t, got, Lte(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Lte for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpLte(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Lte(val), args...)
}

// CmpMap is a shortcut for:
//
//   CmpDeeply(t, got, Map(model, expectedEntries), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Map for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpMap(t TestingT, got interface{}, model interface{}, expectedEntries MapEntries, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Map(model, expectedEntries), args...)
}

// CmpMapEach is a shortcut for:
//
//   CmpDeeply(t, got, MapEach(expectedValue), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#MapEach for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpMapEach(t TestingT, got interface{}, expectedValue interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, MapEach(expectedValue), args...)
}

// CmpN is a shortcut for:
//
//   CmpDeeply(t, got, N(num, tolerance), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#N for details.
//
// N() optional parameter "tolerance" is here mandatory.
// 0 value should be passed to mimic its absence in
// original N() call.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpN(t TestingT, got interface{}, num interface{}, tolerance interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, N(num, tolerance), args...)
}

// CmpNaN is a shortcut for:
//
//   CmpDeeply(t, got, NaN(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NaN for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNaN(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NaN(), args...)
}

// CmpNil is a shortcut for:
//
//   CmpDeeply(t, got, Nil(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Nil for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNil(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Nil(), args...)
}

// CmpNone is a shortcut for:
//
//   CmpDeeply(t, got, None(expectedValues...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#None for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNone(t TestingT, got interface{}, expectedValues []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, None(expectedValues...), args...)
}

// CmpNot is a shortcut for:
//
//   CmpDeeply(t, got, Not(expected), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Not for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNot(t TestingT, got interface{}, expected interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Not(expected), args...)
}

// CmpNotAny is a shortcut for:
//
//   CmpDeeply(t, got, NotAny(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NotAny for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNotAny(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NotAny(expectedItems...), args...)
}

// CmpNotEmpty is a shortcut for:
//
//   CmpDeeply(t, got, NotEmpty(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NotEmpty for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNotEmpty(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NotEmpty(), args...)
}

// CmpNotNaN is a shortcut for:
//
//   CmpDeeply(t, got, NotNaN(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NotNaN for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNotNaN(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NotNaN(), args...)
}

// CmpNotNil is a shortcut for:
//
//   CmpDeeply(t, got, NotNil(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NotNil for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNotNil(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NotNil(), args...)
}

// CmpNotZero is a shortcut for:
//
//   CmpDeeply(t, got, NotZero(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#NotZero for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpNotZero(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, NotZero(), args...)
}

// CmpPPtr is a shortcut for:
//
//   CmpDeeply(t, got, PPtr(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#PPtr for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpPPtr(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, PPtr(val), args...)
}

// CmpPtr is a shortcut for:
//
//   CmpDeeply(t, got, Ptr(val), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Ptr for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpPtr(t TestingT, got interface{}, val interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Ptr(val), args...)
}

// CmpRe is a shortcut for:
//
//   CmpDeeply(t, got, Re(reg, capture), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Re for details.
//
// Re() optional parameter "capture" is here mandatory.
// nil value should be passed to mimic its absence in
// original Re() call.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpRe(t TestingT, got interface{}, reg interface{}, capture interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Re(reg, capture), args...)
}

// CmpReAll is a shortcut for:
//
//   CmpDeeply(t, got, ReAll(reg, capture), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#ReAll for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpReAll(t TestingT, got interface{}, reg interface{}, capture interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, ReAll(reg, capture), args...)
}

// CmpSet is a shortcut for:
//
//   CmpDeeply(t, got, Set(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Set for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSet(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Set(expectedItems...), args...)
}

// CmpShallow is a shortcut for:
//
//   CmpDeeply(t, got, Shallow(expectedPtr), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Shallow for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpShallow(t TestingT, got interface{}, expectedPtr interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Shallow(expectedPtr), args...)
}

// CmpSlice is a shortcut for:
//
//   CmpDeeply(t, got, Slice(model, expectedEntries), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Slice for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSlice(t TestingT, got interface{}, model interface{}, expectedEntries ArrayEntries, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Slice(model, expectedEntries), args...)
}

// CmpSmuggle is a shortcut for:
//
//   CmpDeeply(t, got, Smuggle(fn, expectedValue), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Smuggle for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSmuggle(t TestingT, got interface{}, fn interface{}, expectedValue interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Smuggle(fn, expectedValue), args...)
}

// CmpString is a shortcut for:
//
//   CmpDeeply(t, got, String(expected), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#String for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpString(t TestingT, got interface{}, expected string, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, String(expected), args...)
}

// CmpStruct is a shortcut for:
//
//   CmpDeeply(t, got, Struct(model, expectedFields), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Struct for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpStruct(t TestingT, got interface{}, model interface{}, expectedFields StructFields, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Struct(model, expectedFields), args...)
}

// CmpSubBagOf is a shortcut for:
//
//   CmpDeeply(t, got, SubBagOf(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SubBagOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSubBagOf(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SubBagOf(expectedItems...), args...)
}

// CmpSubMapOf is a shortcut for:
//
//   CmpDeeply(t, got, SubMapOf(model, expectedEntries), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SubMapOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSubMapOf(t TestingT, got interface{}, model interface{}, expectedEntries MapEntries, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SubMapOf(model, expectedEntries), args...)
}

// CmpSubSetOf is a shortcut for:
//
//   CmpDeeply(t, got, SubSetOf(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SubSetOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSubSetOf(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SubSetOf(expectedItems...), args...)
}

// CmpSuperBagOf is a shortcut for:
//
//   CmpDeeply(t, got, SuperBagOf(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SuperBagOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSuperBagOf(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SuperBagOf(expectedItems...), args...)
}

// CmpSuperMapOf is a shortcut for:
//
//   CmpDeeply(t, got, SuperMapOf(model, expectedEntries), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SuperMapOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSuperMapOf(t TestingT, got interface{}, model interface{}, expectedEntries MapEntries, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SuperMapOf(model, expectedEntries), args...)
}

// CmpSuperSetOf is a shortcut for:
//
//   CmpDeeply(t, got, SuperSetOf(expectedItems...), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#SuperSetOf for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpSuperSetOf(t TestingT, got interface{}, expectedItems []interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, SuperSetOf(expectedItems...), args...)
}

// CmpTruncTime is a shortcut for:
//
//   CmpDeeply(t, got, TruncTime(expectedTime, trunc), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#TruncTime for details.
//
// TruncTime() optional parameter "trunc" is here mandatory.
// 0 value should be passed to mimic its absence in
// original TruncTime() call.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpTruncTime(t TestingT, got interface{}, expectedTime interface{}, trunc time.Duration, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, TruncTime(expectedTime, trunc), args...)
}

// CmpZero is a shortcut for:
//
//   CmpDeeply(t, got, Zero(), args...)
//
// See https://godoc.org/github.com/maxatome/go-testdeep#Zero for details.
//
// Returns true if the test is OK, false if it fails.
//
// "args..." are optional and allow to name the test. This name is
// logged as is in case of failure. If len(args) > 1 and the first
// item of args is a string and contains a '%' rune then fmt.Fprintf
// is used to compose the name, else args are passed to fmt.Fprint.
func CmpZero(t TestingT, got interface{}, args ...interface{}) bool {
	// Work around https://github.com/golang/go/issues/26995 issue
	// when corrected, this block should be replaced by t.Helper()
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
	} else {
		t.Helper()
	}

	return CmpDeeply(t, got, Zero(), args...)
}
