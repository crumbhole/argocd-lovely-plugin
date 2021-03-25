package main

import (
	"path/filepath"
)

type PackageDirectories struct {
	dirs []string
}

func (d *PackageDirectories) AddDirectory(path string) {
	if d.dirs == nil {
		d.dirs = make([]string, 0)
	}
	if !d.IsDirectory(path) {
		d.dirs = append(d.dirs, path)
	}
}

func (d *PackageDirectories) IsDirectory(path string) bool {
	for _, knownpath := range d.dirs {
		if path == knownpath {
			return true
		}
	}
	return false
}

func (d *PackageDirectories) KnownSubDirectory(path string) bool {
	for _, dir := range d.dirs {
		if match, _ := filepath.Match(dir+"/*", path); match {
			return true
		}
	}
	return false
}

func (d *PackageDirectories) GetPackages() []string {
	return d.dirs
}
