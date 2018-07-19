package csg

import (
	"testing"
)

func TestSplitPolygon(t *testing.T) {
	pol := newPolygon([]vertex{
		vertex{
			pos:    vector3{-1, 0, -1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{-1, 0, 1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{1, 0, 1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{1, 0, 1},
			normal: vector3{0, 1, 0},
		},
	}, nil)
	pla := plane{normal: vector3{0, 1, 0}, w: 0}
	var coplanarFront, coplanarBack, front, back []polygon
	pla.splitPolygon(pol, &coplanarFront, &coplanarBack, &front, &back)
	if len(coplanarFront) == 0 || len(front)+len(back)+len(coplanarBack) != 0 {
		t.Fail()
	}
}

func TestSplitPolygonSplit(t *testing.T) {
	pol := newPolygon([]vertex{
		vertex{
			pos:    vector3{-1, 0, -1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{-1, 0, 1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{1, 0, 1},
			normal: vector3{0, 1, 0},
		},
		vertex{
			pos:    vector3{1, 0, 1},
			normal: vector3{0, 1, 0},
		},
	}, nil)
	pla := plane{normal: (vector3{0, 1, 1}).unit(), w: 0}
	var coplanarFront, coplanarBack, front, back []polygon
	pla.splitPolygon(pol, &coplanarFront, &coplanarBack, &front, &back)
	t.Logf("coplanarFront: %+v\ncoplanarBack: %+v\nfront: %+v\nback: %+v\n", coplanarFront, coplanarBack, front, back)

	if len(coplanarFront)+len(coplanarBack) != 0 || len(front)+len(back) != 2 {
		t.Fail()
	}
}
