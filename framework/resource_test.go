package framework

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestResources_GetExecutableFile(t *testing.T) {
	fw := NewFramework()
	path := fw.Resource.GetExecutableFile()
	dir := filepath.Dir(path)
	assert.Equal(t, "/tmp", dir)
}

func TestResources_GetExecutablePath(t *testing.T) {
	fw := NewFramework()
	path := fw.Resource.GetExecutablePath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetResourcePath(t *testing.T) {
	fw := NewFramework()
	path := fw.Resource.GetResourcePath("test")
	assert.Equal(t, "/tmp/test", path)
}

func TestResources_GetResourcesPath(t *testing.T) {
	fw := NewFramework()
	path := fw.Resource.GetResourcesPath()
	assert.Equal(t, "/tmp", path)
}

func TestResources_GetUserHomeDirectory(t *testing.T) {
	fw := NewFramework()
	home := fw.Resource.GetUserHomeDirectory()
	assert.Equal(t, "/home/per", home)
}
