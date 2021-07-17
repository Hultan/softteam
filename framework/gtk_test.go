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

func TestGTK_ClearListBox(t *testing.T) {
	gtk.Init(nil)

	listBox, err := gtk.ListBoxNew()
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	label1, err := gtk.LabelNew("label1")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listBox.Add(label1)
	label2, err := gtk.LabelNew("label2")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listBox.Add(label2)
	label3, err := gtk.LabelNew("label3")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	listBox.Add(label3)

	g := NewGTK()
	g.ClearListBox(listBox)
	children := listBox.GetChildren()
	assert.Equal(t, uint(0), children.Length())
}

func TestGTK_ClearFlowBox(t *testing.T) {
	gtk.Init(nil)

	flowBox, err := gtk.FlowBoxNew()
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	label1, err := gtk.LabelNew("label1")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	flowBox.Add(label1)
	label2, err := gtk.LabelNew("label2")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	flowBox.Add(label2)
	label3, err := gtk.LabelNew("label3")
	if err != nil {
		assert.Fail(t, "Should not happen")
		return
	}
	flowBox.Add(label3)

	g := NewGTK()
	g.ClearFlowBox(flowBox)
	children := flowBox.GetChildren()
	assert.Equal(t, uint(0), children.Length())
}
