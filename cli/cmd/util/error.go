package util

import (
	"fmt"
	"os"
)

// FailOnErrorOrForceContinue checks if error occurred and stops the execution
// unless the force flag was specified.
func FailOnErrorOrForceContinue(err error, options *Options) {
	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err.Error())
	if !options.Force {
		os.Exit(1)
	}
}
