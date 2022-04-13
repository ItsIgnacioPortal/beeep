


package main

import (
	"errors"
	"os"
	"syscall"
	"time"
	"unsafe"
)


const (
	
	
	clockTickRate = 1193180

	
	kiocsound = 0x4B2F

	
	evSnd   = 0x12 
	sndTone = 0x02 
)

var (
	
	DefaultFreq = 440.0
	
	DefaultDuration = 200
)


type inputEvent struct {
	Time  syscall.Timeval 
	Type  uint16          
	Code  uint16          
	Value int32           
}


func ioctl(fd, name, data uintptr) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, name, data)
	if e != 0 {
		return e
	}

	return nil
}














func Beep(freq float64, duration int) error {
	if freq == 0 {
		freq = DefaultFreq
	} else if freq > 20000 {
		freq = 20000
	} else if freq < 0 {
		freq = DefaultFreq
	}

	if duration == 0 {
		duration = DefaultDuration
	}

	period := int(float64(clockTickRate) / freq)

	var evdev bool

	f, err := os.OpenFile("/dev/tty0", os.O_WRONLY, 0644)
	if err != nil {
		e := err
		f, err = os.OpenFile("/dev/input/by-path/platform-pcspkr-event-spkr", os.O_WRONLY, 0644)
		if err != nil {
			e = errors.New("beeep: " + e.Error() + "; " + err.Error())

			
			_, err = os.Stdout.Write([]byte{7})
			if err != nil {
				return errors.New(e.Error() + "; " + err.Error())
			}

			return nil
		}

		evdev = true
	}

	defer f.Close()

	if evdev { 
		ev := inputEvent{}
		ev.Type = evSnd
		ev.Code = sndTone
		ev.Value = int32(freq)

		d := *(*[unsafe.Sizeof(ev)]byte)(unsafe.Pointer(&ev))

		
		f.Write(d[:])

		time.Sleep(time.Duration(duration) * time.Millisecond)

		ev.Value = 0
		d = *(*[unsafe.Sizeof(ev)]byte)(unsafe.Pointer(&ev))

		
		f.Write(d[:])
	} else { 
		
		err = ioctl(f.Fd(), kiocsound, uintptr(period))
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(duration) * time.Millisecond)

		
		err = ioctl(f.Fd(), kiocsound, uintptr(0))
		if err != nil {
			return err
		}
	}

	return nil
}
