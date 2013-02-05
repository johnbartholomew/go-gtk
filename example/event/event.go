package main

import (
	"github.com/johnbartholomew/go-gtk/gdk"
	"github.com/johnbartholomew/go-gtk/glib"
	"github.com/johnbartholomew/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", gtk.MainQuit)

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		kev := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		println("key-press-event:", kev.Keyval)
	})
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**gdk.EventMotion)(unsafe.Pointer(&arg))
		println("motion-notify-event:", int(mev.X), int(mev.Y))
	})

	window.SetEvents(int(gdk.POINTER_MOTION_MASK | gdk.POINTER_MOTION_HINT_MASK | gdk.BUTTON_PRESS_MASK))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
