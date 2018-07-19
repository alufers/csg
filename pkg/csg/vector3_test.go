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

func TestCrossProduct(t *testing.T) {
	a := (vector3{1, 2, 3}).cross(vector3{1, 5, 7})
	x := vector3{-1, -4, 3}
	if a != x {
		t.Logf("a = %v", a)
		t.Fail()
	}
}
