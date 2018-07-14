package csg

import (
	"fmt"
	"io"
)

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

func (cs *csgSolid) toObj(out io.Writer) {
	var polyOffsets []int
	var i int
	for _, poly := range cs.polygons {
		polyOffsets = append(polyOffsets, i)
		for _, v := range poly.vertices {
			i++
			fmt.Fprintf(out, "v %v %v %v\nvn %v %v %v\n", v.pos.X, v.pos.Y, v.pos.Z, v.normal.X, v.normal.Y, v.normal.Z)
		}
	}

	for polyIndex, of := range polyOffsets {
		fmt.Fprintf(out, "f ")
		for i := of; i < of+len(cs.polygons[polyIndex].vertices); i++ {
			fmt.Fprintf(out, "%v//%v ", i+1, i+1)
		}
		fmt.Fprintf(out, "\n")
	}
}
