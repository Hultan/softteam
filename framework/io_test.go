package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIO_ExistingFile(t *testing.T) {
	fw := NewFramework()
	exists := fw.IO.FileExists("/home/per/code/softteam/assets/file")
	assert.Equal(t, true, exists)
}

func TestIO_MissingFile(t *testing.T) {
	fw := NewFramework()
	exists := fw.IO.FileExists("/home/per/code/softteam/assets/file2")
	assert.Equal(t, false, exists)
}

func TestIO_ExistingDir(t *testing.T) {
	fw := NewFramework()
	exists := fw.IO.DirectoryExists("/home/per/code/softteam/assets/directory")
	assert.Equal(t, true, exists)
}

func TestIO_MissingDir(t *testing.T) {
	fw := NewFramework()
	exists := fw.IO.FileExists("/home/per/code/softteam/assets/directory2")
	assert.Equal(t, false, exists)
}

func TestIO_DirIsFile(t *testing.T) {
	fw := NewFramework()
	exists := fw.IO.DirectoryExists("/home/per/code/softteam/assets/file")
	assert.Equal(t, false, exists)
}

func TestIO_ReadAllText(t *testing.T) {
	fw := NewFramework()
	text, err := fw.IO.ReadAllText("/home/per/code/softteam/assets/file")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, "this is a test file!\n", text)
}

func TestIO_ReadAllTextFail(t *testing.T) {
	fw := NewFramework()
	text, err := fw.IO.ReadAllText("/home/per/code/softteam/assets/file2")
	if err != nil {
		assert.Equal(t, "", text)
		assert.NotEmpty(t, err)
		return
	}
	assert.Fail(t, "This test should fail!")
}
