package framework

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestResourcesNew(t *testing.T) {
	resource := NewResources()
	assert.NotNil(t, resource)
}

func TestResources_GetExecutableFile(t *testing.T) {
	resource := NewResources()
	path := resource.GetExecutableFile()
	dir := filepath.Dir(path)
	assert.Equal(t, "/tmp", dir)
}

func TestResources_GetExecutablePath(t *testing.T) {
	resource := NewResources()
	path := resource.GetExecutablePath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetResourcePath(t *testing.T) {
	resource := NewResources()
	path := resource.GetResourcePath("test")
	assert.Equal(t, "/tmp/test", path)
}

func TestResources_GetResourcesPath(t *testing.T) {
	resource := NewResources()
	path := resource.GetResourcesPath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetUserHomeDirectory(t *testing.T) {
	resource := NewResources()
	home := resource.GetUserHomeDirectory()
	assert.Equal(t, "/home/per", home)
}
