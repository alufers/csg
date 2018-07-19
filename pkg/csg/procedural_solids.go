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

var cubeData = [][]int{
	{0, 4, 6, 2},
	{1, 3, 7, 5},
	{0, 1, 5, 4},
	{2, 6, 7, 3},
	{0, 2, 3, 1},
	{4, 5, 7, 6},
}

var cubeNormals = []vector3{
	{-1, 0, 0},
	{+1, 0, 0},
	{0, -1, 0},
	{0, +1, 0},
	{0, 0, -1},
	{0, 0, +1},
}

func makeCube(center vector3, size vector3) *csgSolid {
	var polygons []polygon
	for i, data := range cubeData {
		var vertices []vertex
		for _, d := range data {
			vertices = append(vertices, vertex{vector3{
				center.X + size.X*(1*math.Min(1, float64(d&1))-0.5),
				center.Y + size.Y*(1*math.Min(1, float64(d&2))-0.5),
				center.Z + size.Z*(1*math.Min(1, float64(d&4))-0.5),
			}, cubeNormals[i]})
		}
		polygons = append(polygons, newPolygon(vertices, nil))
	}
	return newCsgSolid(polygons)
}
