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
	window.SetTitle("GTK Stock Icons")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)

	store := gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_BOOL, gdkpixbuf.GetType())
	treeview := gtk.NewTreeView()
	swin.Add(treeview)

	treeview.SetModel(store)
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("name", gtk.NewCellRendererText(), "text", 0))
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("check", gtk.NewCellRendererToggle(), "active", 1))
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("icon", gtk.NewCellRendererPixbuf(), "pixbuf", 2))
	n := 0
	gtk.StockListIDs().ForEach(func(d interface{}, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk.TreeIter
		store.Append(&iter)
		store.Set(&iter, id, (n == 1), gtk.NewImage().RenderIcon(id, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
		n = 1 - n
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
