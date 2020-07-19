// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/location"
	"github.com/maxatome/go-testdeep/internal/types"
)

var (
	testDeeper         = reflect.TypeOf((*TestDeep)(nil)).Elem()
	interfaceInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	stringerInterface  = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	errorInterface     = reflect.TypeOf((*error)(nil)).Elem()
	timeType           = reflect.TypeOf(time.Time{})
	intType            = reflect.TypeOf(int(0))
	uint8Type          = reflect.TypeOf(uint8(0))
	runeType           = reflect.TypeOf(rune(0))
	stringType         = reflect.TypeOf("")
	boolType           = reflect.TypeOf(false)
	smuggledGotType    = reflect.TypeOf(SmuggledGot{})
	smuggledGotPtrType = reflect.TypeOf((*SmuggledGot)(nil))
)

// TestingT is the minimal interface used by Cmp to report errors. It
// is commonly implemented by *testing.T and *testing.B.
type TestingT interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Helper()
}

// TestingFT is a deprecated alias of testing.TB. Use testing.TB
// directly in new code.
type TestingFT = testing.TB

// TestDeep is the representation of a go-testdeep operator. It is not
// intended to be used directly, but through Cmp* functions.
type TestDeep interface {
	types.TestDeepStringer
	location.GetLocationer
	// Match checks "got" against the operator. It returns nil if it matches.
	Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error
	setLocation(int)
	// HandleInvalid returns true if the operator is able to handle
	// untyped nil value. Otherwise the untyped nil value is handled
	// generically.
	HandleInvalid() bool
	// TypeBehind returns the type handled by the operator or nil if it
	// is not known. tdhttp helper uses it to know how to unmarshal HTTP
	// responses bodies before comparing them using the operator.
	TypeBehind() reflect.Type
}

// base is a base type providing some methods needed by the TestDeep
// interface.
type base struct {
	types.TestDeepStamp
	location location.Location
}

func pkgFunc(full string) (string, string) {
	// the/package.Foo         → "the/package", "Foo"
	// the/package.(*T).Foo    → "the/package", "(*T).Foo"
	// the/package.glob..func1 → "the/package", "glob..func1"
	sp := strings.LastIndexByte(full, '/')
	if sp < 0 {
		sp = 0 // std package without any '/' in name
	}

	dp := strings.IndexByte(full[sp:], '.')
	if dp < 0 {
		return full, ""
	}
	dp += sp
	return full[:dp], full[dp+1:]
}

func (t *base) setLocation(callDepth int) {
	var ok bool
	t.location, ok = location.New(callDepth)
	if !ok {
		t.location.File = "???"
		t.location.Line = 0
		return
	}

	// Here package is github.com/maxatome/go-testdeep, or its vendored
	// counterpart
	var pkg string
	pkg, t.location.Func = pkgFunc(t.location.Func)

	// Try to go one level upper, if we are still in go-testdeep package
	cmpLoc, ok := location.New(callDepth + 1)
	if ok {
		cmpPkg, _ := pkgFunc(cmpLoc.Func)
		if cmpPkg == pkg {
			t.location.File = cmpLoc.File
			t.location.Line = cmpLoc.Line
			t.location.BehindCmp = true
		}
	}
}

// GetLocation returns a copy of the location.Location where the TestDeep
// operator has been created.
func (t *base) GetLocation() location.Location {
	return t.location
}

// HandleInvalid tells go-testdeep internals that this operator does
// not handle nil values directly.
func (t base) HandleInvalid() bool {
	return false
}

// TypeBehind returns the type handled by the operator. Only few operators
// knows the type they are handling. If they do not know, nil is
// returned.
func (t base) TypeBehind() reflect.Type {
	return nil
}

// newBase returns a new base struct with location.Location set to the
// "callDepth" depth.
func newBase(callDepth int) (b base) {
	b.setLocation(callDepth)
	return
}

// baseOKNil is a base type providing some methods needed by the TestDeep
// interface, for operators handling nil values.
type baseOKNil struct {
	base
}

// HandleInvalid tells go-testdeep internals that this operator
// handles nil values directly.
func (t baseOKNil) HandleInvalid() bool {
	return true
}

// newBaseOKNil returns a new baseOKNil struct with location.Location set to
// the "callDepth" depth.
func newBaseOKNil(callDepth int) (b baseOKNil) {
	b.setLocation(callDepth)
	return
}