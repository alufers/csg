package csg

type polygon struct {
	vertices []vertex
	shared   interface{}
	plane    plane
}

func newPolygon(vertices []vertex, shared interface{}) polygon {
	return polygon{
		vertices: vertices,
		shared:   shared,
		plane:    planeFromPoints(vertices[0].pos, vertices[1].pos, vertices[2].pos),
	}
}

func (p polygon) flipped() polygon {
	l := len(p.vertices)
	verts := make([]vertex, l)
	for i, v := range p.vertices {
		verts[l-i-1] = v.flipped()
	}
	return polygon{
		plane:    p.plane.flipped(),
		vertices: verts,
		shared:   p.shared,
	}
}
