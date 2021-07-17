package framework

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestResources_New(t *testing.T) {
	resource := NewResource()
	assert.NotNil(t, resource)
}

func TestResources_GetExecutableFile(t *testing.T) {
	resource := NewResource()
	path := resource.GetExecutableFile()
	dir := filepath.Dir(path)
	assert.Equal(t, "/tmp", dir)
}

func TestResources_GetExecutablePath(t *testing.T) {
	resource := NewResource()
	path := resource.GetExecutablePath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetResourcePath(t *testing.T) {
	resource := NewResource()
	path := resource.GetResourcePath("test")
	assert.Equal(t, "/tmp/test", path)
}

func TestResources_GetResourcesPath(t *testing.T) {
	resource := NewResource()
	path := resource.GetResourcesPath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetUserHomeDirectory(t *testing.T) {
	resource := NewResource()
	home := resource.GetUserHomeDirectory()
	assert.Equal(t, "/home/per", home)
}
