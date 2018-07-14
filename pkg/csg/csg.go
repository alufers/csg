package csg

import (
	"fmt"
)

func Run() {
	//sp1 := makeSphere(vector3{0, 0, 0}, 2)
	//sp2 := makeSphere(vector3{0, 0, 0}, 1)
	//p1.intersect(sp2)
	//.toObj(os.Stdout)

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
	fmt.Printf("coplanarFront: %+v\ncoplanarBack: %+v\nfront: %+v\nback: %+v\n", coplanarFront, coplanarBack, front, back)
}
