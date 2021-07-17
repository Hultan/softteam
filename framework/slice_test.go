package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_ContainsInt(t *testing.T) {
	fw := NewFramework()
	data := []int{23,12,45,26,74}
	result := fw.Slice.ContainsInt(data, 23)
	assert.True(t, result)
	result = fw.Slice.ContainsInt(data, 24)
	assert.False(t, result)
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
	data := []int{23,12,45,26,74}
	data = fw.Slice.RemoveIntAt(data, 0)
	result := fw.Slice.ContainsInt(data, 23)
	assert.False(t, result)
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
	data := []int{23,12,45,26,74}
	data = fw.Slice.RemoveIntAtFast(data, 0)
	result := fw.Slice.ContainsInt(data, 23)
	assert.False(t, result)
}

func TestSlice_RemoveStringAtFast(t *testing.T) {
	fw := NewFramework()
	data := []string{"per", "fredrik", "joakim"}
	data = fw.Slice.RemoveStringAtFast(data, 1)
	result := fw.Slice.ContainsString(data, "fredrik")
	assert.False(t, result)
}