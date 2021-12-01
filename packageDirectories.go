package main

import (
	"path/filepath"
)

// PackageDirectories is an array of sub-application paths
type PackageDirectories struct {
	dirs []string
}

// AddDirectory Adds a directory to PackageDirectories if it isn't in there already
func (d *PackageDirectories) AddDirectory(path string) {
	if d.dirs == nil {
		d.dirs = make([]string, 0)
	}
	if !d.IsDirectory(path) {
		d.dirs = append(d.dirs, path)
	}
}

// IsDirectory returns true if a path is in PackageDirectories
func (d *PackageDirectories) IsDirectory(path string) bool {
	for _, knownpath := range d.dirs {
		if path == knownpath {
			return true
		}
	}
	return false
}

// KnownSubDirectory returns true if a path is in PackageDirectories or is a subdirectory of one that is
func (d *PackageDirectories) KnownSubDirectory(path string) bool {
	if d.IsDirectory(path) {
		return true
	}
	for _, dir := range d.dirs {
		if match, _ := filepath.Match(dir+"/*", path); match {
			return true
		}
	}
	return false
}

// GetPackages returns an array of paths that make up this PackageDirectories
func (d *PackageDirectories) GetPackages() []string {
	return d.dirs
}
