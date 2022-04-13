


package main

import (
	"syscall"
)

var (
	
	DefaultFreq = 587.0
	
	DefaultDuration = 500
)


func Beep(freq float64, duration int) error {
	if freq == 0 {
		freq = DefaultFreq
	} else if freq > 32767 {
		freq = 32767
	} else if freq < 37 {
		freq = DefaultFreq
	}

	if duration == 0 {
		duration = DefaultDuration
	}

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	beep32, _ := syscall.GetProcAddress(kernel32, "Beep")

	defer syscall.FreeLibrary(kernel32)

	_, _, e := syscall.Syscall(uintptr(beep32), uintptr(2), uintptr(int(freq)), uintptr(duration), 0)
	if e != 0 {
		return e
	}

	return nil
}
