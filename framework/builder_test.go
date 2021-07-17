package framework

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder_New(t *testing.T) {
	gtk.Init(nil)

	builder, err := NewBuilder("/home/per/code/softteam/assets/main.glade")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotNil(t, builder)
}

func TestBuilder_GetObject(t *testing.T) {
	gtk.Init(nil)

	builder, err := NewBuilder("/home/per/code/softteam/assets/main.glade")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	window := builder.GetObject("mainWindow").(*gtk.Window)
	assert.NotNil(t, window)
	name, err := window.GetName()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, "Main Window", name)
}
