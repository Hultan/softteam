package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess_OpenEcho(t *testing.T) {
	fw := NewFramework()
	output := fw.Process.Open("echo", "hello")
	assert.Equal(t, "hello\n", output)
}
