


package main

import (
	"syscall/js"
)





func Notify(title, message, appIcon string) (err error) {
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

	n := js.Global().Get("Notification")

	opts := js.Global().Get("Object").Invoke()
	opts.Set("body", message)
	opts.Set("icon", pathAbs(appIcon))

	if n.Get("permission").String() == "granted" {
		n.New(js.ValueOf(title), opts)
	} else {
		var f js.Func
		f = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if args[0].String() == "granted" {
				n.New(js.ValueOf(title), opts)
			}
			f.Release()
			return nil
		})

		n.Call("requestPermission", f)
	}

	return
}
