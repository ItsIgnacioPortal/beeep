


package main

import (
	toast "github.com/go-toast/toast"
)


func Alert(title, message, appIcon string) error {
	if isWindows10 {
		note := toastNotification(title, message, pathAbs(appIcon))
		note.Audio = toast.Default
		return note.Push()
	}

	if err := Notify(title, message, appIcon); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
