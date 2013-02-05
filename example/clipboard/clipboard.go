package main

import (
	"github.com/johnbartholomew/go-gtk/gdk"
	"github.com/johnbartholomew/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	clipboard := gtk.NewClipboardGetForDisplay(
		gdk.DisplayGetDefault(),
		gdk.SELECTION_CLIPBOARD)
	println(clipboard.WaitForText())
	clipboard.SetText("helloworld")
	gtk.MainIterationDo(true)
	clipboard.Store()
	gtk.MainIterationDo(true)
}
