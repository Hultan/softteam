package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProcess(t *testing.T) {
	process := NewProcess()
	assert.NotNil(t, process)
}
