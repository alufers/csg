package csg

import (
	"testing"
)

func TestPlus(t *testing.T) {
	a := (vector3{1, 2, 3}).plus(vector3{9, 15, 33})
	if a.X != 10 || a.Y != 17 || a.Z != 36 {
		t.Fail()
	}
}
