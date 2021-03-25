package main

import (
	"testing"
)

func TestNotDir(t *testing.T) {
	p := PackageDirectories{}
	if p.IsDirectory(`foo`) {
		t.Errorf("Directory foo IsDirectory when it isn't")
	}
}

func TestIsDir(t *testing.T) {
	p := PackageDirectories{}
	p.AddDirectory(`foo`)
	p.AddDirectory(`blueberry`)
	if !p.IsDirectory(`foo`) {
		t.Errorf("Directory foo !IsDirectory when it is")
	}
	if p.IsDirectory(`bar`) {
		t.Errorf("Directory bar IsDirectory when it isn't")
	}
}
