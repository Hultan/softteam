package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_ContainsInt(t *testing.T) {
	fw := NewFramework()
	data := []int{23, 12, 45, 26, 74}
	ok := fw.Slice.ContainsInt(data, 12)
	assert.True(t, ok)
	ok = fw.Slice.ContainsInt(data, 24)
	assert.False(t, ok)
}

func TestSlice_ContainsString(t *testing.T) {
	fw := NewFramework()
	data := []string{"per", "fredrik", "joakim"}
	result := fw.Slice.ContainsString(data, "joakim")
	assert.True(t, result)
	result = fw.Slice.ContainsString(data, "apple")
	assert.False(t, result)
}

func TestSlice_RemoveIntAt(t *testing.T) {
	fw := NewFramework()
	data := []int{23, 12, 45, 26, 74}
	data = fw.Slice.RemoveIntAt(data, 0)
	ok := fw.Slice.ContainsInt(data, 23)
	assert.False(t, ok)
}

func TestSlice_RemoveStringAt(t *testing.T) {
	fw := NewFramework()
	data := []string{"per", "fredrik", "joakim"}
	data = fw.Slice.RemoveStringAt(data, 1)
	result := fw.Slice.ContainsString(data, "fredrik")
	assert.False(t, result)
}

func TestSlice_RemoveIntAtFast(t *testing.T) {
	fw := NewFramework()
	data := []int{23, 12, 45, 26, 74}
	data = fw.Slice.RemoveIntAtFast(data, 0)
	ok := fw.Slice.ContainsInt(data, 23)
	assert.False(t, ok)
}

func TestSlice_RemoveStringAtFast(t *testing.T) {
	fw := NewFramework()
	data := []string{"per", "fredrik", "joakim"}
	data = fw.Slice.RemoveStringAtFast(data, 1)
	result := fw.Slice.ContainsString(data, "fredrik")
	assert.False(t, result)
}

func TestSlice_IndexOfInt(t *testing.T) {
	fw := NewFramework()
	data := []int{23, 12, 45, 26, 74}
	index := fw.Slice.IndexOfInt(data, 12)
	assert.Equal(t, 1, index)
	index = fw.Slice.IndexOfInt(data, 24)
	assert.Equal(t, -1, index)
}

func TestSlice_IndexOfString(t *testing.T) {
	fw := NewFramework()
	data := []string{"per", "fredrik", "joakim"}
	result := fw.Slice.IndexOfString(data, "joakim")
	assert.Equal(t, 2, result)
	result = fw.Slice.IndexOfString(data, "apple")
	assert.Equal(t, -1, result)
}
