package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_New(t *testing.T) {
	slice := NewSlice()
	assert.NotNil(t, slice)
}

func TestSlice_ContainsInt(t *testing.T) {
	slice := NewSlice()
	data := []int{23,12,45,26,74}
	result := slice.ContainsInt(data, 23)
	assert.True(t, result)
	result = slice.ContainsInt(data, 24)
	assert.False(t, result)
}

func TestSlice_ContainsString(t *testing.T) {
	slice := NewSlice()
	data := []string{"per", "fredrik", "joakim"}
	result := slice.ContainsString(data, "joakim")
	assert.True(t, result)
	result = slice.ContainsString(data, "apple")
	assert.False(t, result)
}

func TestSlice_RemoveIntAt(t *testing.T) {
	slice := NewSlice()
	data := []int{23,12,45,26,74}
	data = slice.RemoveIntAt(data, 0)
	result := slice.ContainsInt(data, 23)
	assert.False(t, result)
}

func TestSlice_RemoveStringAt(t *testing.T) {
	slice := NewSlice()
	data := []string{"per", "fredrik", "joakim"}
	data = slice.RemoveStringAt(data, 1)
	result := slice.ContainsString(data, "fredrik")
	assert.False(t, result)
}

func TestSlice_RemoveIntAtFast(t *testing.T) {
	slice := NewSlice()
	data := []int{23,12,45,26,74}
	data = slice.RemoveIntAtFast(data, 0)
	result := slice.ContainsInt(data, 23)
	assert.False(t, result)
}

func TestSlice_RemoveStringAtFast(t *testing.T) {
	slice := NewSlice()
	data := []string{"per", "fredrik", "joakim"}
	data = slice.RemoveStringAtFast(data, 1)
	result := slice.ContainsString(data, "fredrik")
	assert.False(t, result)
}