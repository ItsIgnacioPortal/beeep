
package main

import (
	"errors"
	"path/filepath"
	"runtime"
)

var (
	
	ErrUnsupported = errors.New("beeep: unsupported operating system: " + runtime.GOOS)
)

func pathAbs(path string) string {
	var err error
	var abs string

	if path != "" {
		abs, err = filepath.Abs(path)
		if err != nil {
			abs = path
		}
	}

	return abs
}
