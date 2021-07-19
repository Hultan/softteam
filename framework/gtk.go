package framework

import (
	"github.com/gotk3/gotk3/gtk"
)

type Gtk struct {
}

// ClearListBox : Clears a gtk.ListBox
func (g *Gtk) ClearListBox(list *gtk.ListBox) {
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

// ClearFlowBox : Clears a gtk.FlowBox
func (g *Gtk) ClearFlowBox(list *gtk.FlowBox) {
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

// CreateBuilder : Creates the actual gtk.Builder
func (g *Gtk) CreateBuilder(fileNameOrPath string) (*GtkBuilder, error) {
	fw := NewFramework()
	if !fw.IO.FileExists(fileNameOrPath) {
		fileNameOrPath = fw.Resource.GetResourcePath(fileNameOrPath)
	}

	builder, err := gtk.BuilderNewFromFile(fileNameOrPath)
	if err != nil {
		return nil, err
	}

	return &GtkBuilder{builder}, nil
}
