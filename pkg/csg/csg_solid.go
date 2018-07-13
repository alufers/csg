package csg

type csgSolid struct {
	polygons []polygon
}

func newCsgSolid(polygons []polygon) *csgSolid {
	return &csgSolid{
		polygons: polygons,
	}
}

func (cs *csgSolid) clone() *csgSolid {
	newSolid := &csgSolid{}
	newSolid.polygons = make([]polygon, len(cs.polygons))
	copy(newSolid.polygons, cs.polygons)
	return newSolid
}

func (cs *csgSolid) subtract(other *csgSolid) *csgSolid {
	a := newCsgNode(cs.clone().polygons)
	b := newCsgNode(other.clone().polygons)
	a.invert()
	a.clipTo(b)
	b.clipTo(a)
	b.invert()
	b.clipTo(a)
	b.invert()
	a.build(b.allPolygons())
	a.invert()
	return newCsgSolid(a.allPolygons())
}

func (cs *csgSolid) union(other *csgSolid) *csgSolid {
	a := newCsgNode(cs.clone().polygons)
	b := newCsgNode(other.clone().polygons)
	a.clipTo(b)
	b.clipTo(a)
	b.invert()
	b.clipTo(a)
	b.invert()
	a.build(b.allPolygons())
	return newCsgSolid(a.allPolygons())
}

func (cs *csgSolid) intersect(other *csgSolid) *csgSolid {
	a := newCsgNode(cs.clone().polygons)
	b := newCsgNode(other.clone().polygons)
	a.invert()
	b.clipTo(a)
	b.invert()
	a.clipTo(b)
	b.clipTo(a)
	a.build(b.allPolygons())
	a.invert()
	return newCsgSolid(a.allPolygons())
}
