package framework

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// Resource : Handles SoftTeam resources
type Resource struct {
}

// GetExecutableFile : Returns the path of the executable file
func (r *Resource) GetExecutableFile() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return ex
}

// GetExecutablePath : Returns the directory of the executable file
func (r *Resource) GetExecutablePath() string {
	return filepath.Dir(r.GetExecutableFile())
}

// GetResourcePath : Gets the path for a single resource file
func (r *Resource) GetResourcePath(fileName string) string {
	resourcesPath := r.GetResourcesPath()
	resourcePath := path.Join(resourcesPath, fileName)
	return resourcePath
}

func (r *Resource) GetUserHomeDirectory() string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if home != "" {
		return home
	}

	return os.Getenv("HOME")
}

// GetResourcesPath : Returns the resources path
func (r *Resource) GetResourcesPath() string {
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

// checkPathsExists : Returns the first path that exists
func (r *Resource) checkPathsExists(pathsToCheck []string) (string, error) {
	fw := NewFramework()
	for _, pathToCheck := range pathsToCheck {
		if fw.IO.DirectoryExists(pathToCheck) {
			return pathToCheck, nil
		}
	}
	return "", errors.New("paths do not exist")
}
