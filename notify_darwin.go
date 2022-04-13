


package main

import (
	"os/exec"
)




func Notify(title, message, appIcon string) error {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return err
	}

	cmd := exec.Command(osa, "-e", `display notification "`+message+`" with title "`+title+`"`)
	return cmd.Run()
}
