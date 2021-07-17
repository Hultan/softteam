package framework

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGTK(t *testing.T) {
	gtk := NewGTK()
	assert.NotNil(t, gtk)
}

func TestGTK_ClearList(t *testing.T) {
	gtk.Init(nil)

	listbox, err := gtk.ListBoxNew()
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	label1, err := gtk.LabelNew("label1")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listbox.Add(label1)
	label2, err := gtk.LabelNew("label2")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listbox.Add(label2)
	label3, err := gtk.LabelNew("label3")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listbox.Add(label3)

	g := NewGTK()
	g.ClearListBox(listbox)
	children := listbox.GetChildren()
	assert.Equal(t, uint(0), children.Length())
}
