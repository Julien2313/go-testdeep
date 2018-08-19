// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"os"
	"strconv"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
)

// ContextConfig allows to configure finely how tests failures are rendered.
//
// See NewT function to use it.
type ContextConfig struct {
	// RootName is the string used to represent the root of got data. It
	// defaults to "DATA". For an HTTP response body, it could be "BODY"
	// for example.
	RootName string
	// MaxErrors is the maximal number of errors to dump in case of Cmp*
	// failure.
	//
	// It defaults to 10 except if the environment variable
	// TESTDEEP_MAX_ERRORS is set. In this latter case, the
	// TESTDEEP_MAX_ERRORS value is converted to an int and used as is.
	//
	// Setting it to 0 has the same effect as 1: only the first error
	// will be dumped without the "Too many errors" error.
	//
	// Setting it to a negative number means no limit: all errors
	// will be dumped.
	MaxErrors int
	// FailureIsFatal allows to Fatal() (instead of Error()) when a test
	// fails. Using *testing.T instance as
	// t.TestingFT value, FailNow() is called behind the scenes when
	// Fatal() is called. See testing documentation for details.
	FailureIsFatal bool
}

const (
	contextDefaultRootName = "DATA"
	contextPanicRootName   = "FUNCTION"
	envMaxErrors           = "TESTDEEP_MAX_ERRORS"
)

func getMaxErrorsFromEnv() int {
	env := os.Getenv(envMaxErrors)
	if env != "" {
		n, err := strconv.Atoi(env)
		if err == nil {
			return n
		}
	}
	return 10
}

// DefaultContextConfig is the default configuration used to render
// tests failures. If overridden, new settings will impact all Cmp*
// functions and *T methods (if not specifically configured.)
var DefaultContextConfig = ContextConfig{
	RootName:       contextDefaultRootName,
	MaxErrors:      getMaxErrorsFromEnv(),
	FailureIsFatal: false,
}

func (c *ContextConfig) sanitize() {
	if c.RootName == "" {
		c.RootName = DefaultContextConfig.RootName
	}
	if c.MaxErrors == 0 {
		c.MaxErrors = DefaultContextConfig.MaxErrors
	}
}

// newContext creates a new ctxerr.Context using DefaultContextConfig
// configuration.
func newContext() ctxerr.Context {
	return newContextWithConfig(DefaultContextConfig)
}

// newContextWithConfig creates a new ctxerr.Context using a specific
// configuration.
func newContextWithConfig(config ContextConfig) (ctx ctxerr.Context) {
	config.sanitize()

	ctx = ctxerr.Context{
		Path:           config.RootName,
		Visited:        map[ctxerr.Visit]bool{},
		MaxErrors:      config.MaxErrors,
		FailureIsFatal: config.FailureIsFatal,
	}

	ctx.InitErrors()
	return
}

// newBooleanContext creates a new boolean ctxerr.Context.
func newBooleanContext() ctxerr.Context {
	return ctxerr.Context{
		Visited:      map[ctxerr.Visit]bool{},
		BooleanError: true,
	}
}
