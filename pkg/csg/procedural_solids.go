package csg

import (
	"math"
)

func makeSphere(center vector3, radius float64) *csgSolid {
	const (
		slices = 16
		stacks = 8
	)
	var polygons []polygon
	var vertices []vertex
	makeVertex := func(theta, phi float64) {
		theta *= math.Pi * 2
		phi *= math.Pi
		dir := vector3{
			math.Cos(theta) * math.Sin(phi),
			math.Cos(phi),
			math.Sin(theta) * math.Sin(phi),
		}
		vertices = append(vertices, vertex{
			pos:    dir.times(radius).plus(center),
			normal: dir,
		})

	}
	for i := 0; i < slices; i++ {
		for j := 0; j < stacks; j++ {
			vertices = []vertex{}
			makeVertex(float64(i)/float64(slices), float64(j)/float64(stacks))
			if j > 0 {
				makeVertex(float64(i+1)/float64(slices), float64(j)/float64(stacks))
			}
			if j < stacks-1 {
				makeVertex(float64(i+1)/float64(slices), float64(j+1)/float64(stacks))
			}

			makeVertex(float64(i)/float64(slices), float64(j+1)/float64(stacks))
			polygons = append(polygons, newPolygon(vertices, nil))
		}
	}
	return newCsgSolid(polygons)
}
