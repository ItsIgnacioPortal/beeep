


package main

import (
	"syscall/js"
)

var (
	
	DefaultFreq = 0.0
	
	DefaultDuration = 0
)


func Beep(freq float64, duration int) (err error) {
	defer func() {
		e := recover()

		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	a := js.Global().Get("document").Call("createElement", "audio")
	a.Set("src", `data:audio/wav;base64,
	a.Call("play")

	return
}
