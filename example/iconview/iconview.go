package main

import (
	"github.com/johnbartholomew/go-gtk/gdkpixbuf"
	"github.com/johnbartholomew/go-gtk/glib"
	"github.com/johnbartholomew/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Icon View")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)

	store := gtk.NewListStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING)
	iconview := gtk.NewIconViewWithModel(store)
	iconview.SetPixbufColumn(0)
	iconview.SetTextColumn(1)
	swin.Add(iconview)

	gtk.StockListIDs().ForEach(func(d interface{}, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk.TreeIter
		store.Append(&iter)
		store.Set(&iter,
			gtk.NewImage().RenderIcon(id, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf,
			id)
	})

	window.Add(swin)
	window.SetSizeRequest(500, 200)
	window.ShowAll()

	gtk.Main()
}
