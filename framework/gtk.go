package framework

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type GTK struct {
	Builder *gtk.Builder
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

// ClearFlowBox : Clears a gtk.FlowBox
func (g *GTK) ClearFlowBox(list *gtk.FlowBox) {
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
func (g *GTK) CreateBuilder(fileNameOrPath string) (*gtk.Builder, error) {
	fw := NewFramework()
	if !fw.IO.FileExists(fileNameOrPath) {
		fileNameOrPath = fw.Resource.GetResourcePath(fileNameOrPath)
	}

	builder, err := gtk.BuilderNewFromFile(fileNameOrPath)
	if err != nil {
		return nil, err
	}

	g.Builder = builder
	return builder, nil
}

// GetObject : Gets a gtk object by name
func (g *GTK) GetObject(name string) glib.IObject {
	if g.Builder == nil {
		panic(errors.New("need to call CreateBuilder first"))
	}
	obj, err := g.Builder.GetObject(name)
	if err != nil {
		panic(err)
	}

	return obj
}
