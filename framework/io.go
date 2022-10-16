package framework

import (
	"io"
	"os"
)

type IO struct {
}

// FileExists : Checks if the specified file exists
func (i *IO) FileExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return !info.IsDir()
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// DirectoryExists : Checks if the specified directory exists
func (i *IO) DirectoryExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return info.IsDir()
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// ReadAllText : Reads a text file and returns the content
func (i *IO) ReadAllText(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(file)

	if err = file.Close(); err != nil {
		return "", err
	}

	return string(bytes), nil
}
