//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!js

package main

// Notify sends desktop notification.
func Notify(title, message string) error {
	return ErrUnsupported
}
