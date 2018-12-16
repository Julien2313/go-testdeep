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
)

type tdAll struct {
	tdList
}

var _ TestDeep = &tdAll{}

// All operator compares data against several expected values. During
// a match, all of them have to match to succeed.
//go:noinline
func All(expectedValues ...interface{}) TestDeep {
	return &tdAll{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAll) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	var origErr *ctxerr.Error
	for idx, item := range a.items {
		// Use deepValueEqualFinal here instead of deepValueEqual as we
		// want to know whether an error occurred or not, we do not want
		// to accumulate it silently
		origErr = deepValueEqualFinal(
			ctx.ResetErrors().
				AddDepth(fmt.Sprintf("<All#%d/%d>", idx+1, len(a.items))),
			got, item)
		if origErr != nil {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			err := &ctxerr.Error{
				Message:  fmt.Sprintf("compared (part %d of %d)", idx+1, len(a.items)),
				Got:      got,
				Expected: item,
			}
			if item.IsValid() && item.Type().Implements(testDeeper) {
				err.Origin = origErr
			}
			return ctx.CollectError(err)
		}
	}
	return nil
}
