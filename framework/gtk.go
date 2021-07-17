package framework

import "github.com/gotk3/gotk3/gtk"

type GTK struct {

}

// NewGTK : Creates a new GTK type
func NewGTK() *GTK{
	return new(GTK)
}

// ClearListBox : Clears a gtk.ListBox
func (g *GTK) ClearListBox(list *gtk.ListBox) {
	children := list.GetChildren()
	if children == nil {
		return
	}
	var i uint = 0
	for ; i < children.Length(); {
		widget, _ := children.NthData(i).(*gtk.Widget)
		list.Remove(widget)
		i++
	}
}
