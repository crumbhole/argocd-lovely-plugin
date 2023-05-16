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
	p.AddDirectory(`fish`)
	p.AddDirectory(`fish/haddock`)
	if !p.IsDirectory(`foo`) {
		t.Errorf("Directory foo !IsDirectory when it is")
	}
	if !p.IsDirectory(`blueberry`) {
		t.Errorf("Directory blueberry !IsDirectory when it is")
	}
	if !p.IsDirectory(`fish`) {
		t.Errorf("Directory fish !IsDirectory when it is")
	}
	if !p.IsDirectory(`fish/haddock`) {
		t.Errorf("Directory fish/haddock !IsDirectory when it is")
	}
	if p.IsDirectory(`bar`) {
		t.Errorf("Directory bar IsDirectory when it isn't")
	}
}

func TestNotKnownSubDirectory(t *testing.T) {
	p := PackageDirectories{}
	if p.KnownSubDirectory(`foo`) {
		t.Errorf("Directory foo KnownSubDirectory when it isn't")
	}
}

func TestKnownSubDirectory(t *testing.T) {
	p := PackageDirectories{}
	p.AddDirectory(`foo`)
	p.AddDirectory(`blueberry`)
	p.AddDirectory(`fish`)
	p.AddDirectory(`fish/haddock`)
	if !p.KnownSubDirectory(`foo`) {
		t.Errorf("Directory foo !KnownSubDirectory when it is")
	}
	if !p.KnownSubDirectory(`blueberry/blue`) {
		t.Errorf("Directory blueberry/blue !KnownSubDirectory when it is")
	}
	if !p.KnownSubDirectory(`fish`) {
		t.Errorf("Directory fish !KnownSubDirectory when it is")
	}
	if !p.KnownSubDirectory(`fish/skate`) {
		t.Errorf("Directory fish/skate !KnownSubDirectory when it is")
	}
	if p.KnownSubDirectory(`bar`) {
		t.Errorf("Directory bar KnownSubDirectory when it isn't")
	}
}
