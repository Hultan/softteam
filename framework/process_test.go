package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess_OpenEcho(t *testing.T) {
	fw := NewFramework()
	output := fw.Process.Open("echo", "hello")
	assert.Equal(t, "hello\n", output)
}
