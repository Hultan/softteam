package framework

import (
	"io/ioutil"
	"os"
)

type IO struct {

}

func NewIO() *IO {
	return new(IO)
}

func (i *IO) FileExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return !info.IsDir()
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func (i *IO) DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (i *IO) ReadAllText(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(file)

	if err = file.Close(); err != nil {
		return "", err
	}

	return string(bytes), nil
}


