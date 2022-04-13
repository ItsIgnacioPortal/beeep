


package main

import (
	"os"
	"os/exec"
)

var (
	
	DefaultFreq = 0.0
	
	DefaultDuration = 0
)


func Beep(freq float64, duration int) error {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		
		_, err = os.Stdout.Write([]byte{7})
		return err
	}

	cmd := exec.Command(osa, "-e", `beep`)
	return cmd.Run()
}
