package main

import (
	"github.com/johnbartholomew/go-gtk/example/builder/callback"
	"github.com/johnbartholomew/go-gtk/gtk"
	"os"
)

//"github.com/johnbartholomew/go-gtk/example/builder/callback"
func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("hello.ui")
	builder.ConnectSignals(nil)
	obj := builder.GetObject("window1")

	callback.Init(builder)

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk.MainQuit)

	gtk.Main()
}
