package framework

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// Resources : Handles SoftTeam resources
type Resources struct {
}

func NewResources() *Resources {
	return new(Resources)
}

// GetExecutableFile : Returns the path of the executable file
func (r *Resources) GetExecutableFile() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return ex
}

// GetExecutablePath : Returns the directory of the executable file
func (r *Resources) GetExecutablePath() string {
	return filepath.Dir(r.GetExecutableFile())
}

// GetResourcePath : Gets the path for a single resource file
func (r *Resources) GetResourcePath(fileName string) string {
	resourcesPath := r.GetResourcesPath()
	resourcePath := path.Join(resourcesPath, fileName)
	return resourcePath
}

func (r *Resources) GetUserHomeDirectory() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

// GetResourcesPath : Returns the resources path
func (r *Resources) GetResourcesPath() string {
	executablePath := r.GetExecutablePath()

	var pathsToCheck []string
	pathsToCheck = append(pathsToCheck, path.Join(executablePath, "../assets"))
	pathsToCheck = append(pathsToCheck, path.Join(executablePath, "assets"))

	dir, err := r.checkPathsExists(pathsToCheck)
	if err != nil {
		return executablePath
	}
	return dir
}

func (r *Resources) checkPathsExists(pathsToCheck []string) (string, error) {
	io := NewIO()
	for _, pathToCheck := range pathsToCheck {
		if io.DirectoryExists(pathToCheck) {
			return pathToCheck, nil
		}
	}
	return "", errors.New("paths do not exist")
}
