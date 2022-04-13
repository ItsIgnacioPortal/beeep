## beeep
[![TravisCI Build Status](https:
[![AppVeyor Build Status](https:
[![GoDoc](https:
[![Go Report Card](https:
<!--[![Go Cover](http:

`beeep` provides a cross-platform library for sending desktop notifications, alerts and beeps.

### Installation

    go get -u github.com/gen2brain/beeep

### Build tags

* `nodbus` - disable `godbus/dbus` and use only `notify-send`

### Examples

```go
err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
if err != nil {
    panic(err)
}
```

```go
err := beeep.Notify("Title", "Message body", "assets/information.png")
if err != nil {
    panic(err)
}
```

```go
err := beeep.Alert("Title", "Message body", "assets/warning.png")
if err != nil {
    panic(err)
}
```


## More

For cross-platform dialogs and input boxes see [dlgs](https:
