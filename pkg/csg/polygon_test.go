package csg

import (
	"testing"
)

func TestPolygonFlip(t *testing.T) {
	poly := newPolygon([]vertex{
		{pos: vector3{1, 1, 0}},
		{pos: vector3{2, 1, 0}},
		{pos: vector3{1, 2, 0}},
	}, nil)
	if (poly.plane.normal != vector3{Z: 1}) {
		t.Fatal("Bad normal before flip")
	}
	flipped := poly.flipped()
	if (flipped.plane.normal != vector3{Z: -1}) {
		t.Fatal("Bad normal after flip")
	}
}
