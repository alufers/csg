package csg

type csgNode struct {
	front    *csgNode
	back     *csgNode
	polygons []polygon
	plane    *plane
}

func newCsgNode(polygons []polygon) *csgNode {
	n := &csgNode{}
	n.build(polygons)
	return n
}

func (cn *csgNode) clone() *csgNode {
	newNode := &csgNode{}
	if cn.front != nil {
		newNode.front = cn.front.clone()
	}
	if cn.back != nil {
		newNode.back = cn.back.clone()
	}
	if cn.plane != nil {
		*newNode.plane = *cn.plane
	}
	newNode.polygons = make([]polygon, len(cn.polygons))
	copy(newNode.polygons, cn.polygons)
	return newNode
}

func (cn *csgNode) invert() {
	for i := range cn.polygons {
		cn.polygons[i] = cn.polygons[i].flipped()
	}
	*cn.plane = cn.plane.flipped()
	if cn.front != nil {
		cn.front.invert()
	}
	if cn.back != nil {
		cn.back.invert()
	}
	cn.back, cn.front = cn.front, cn.back
}

func (cn *csgNode) allPolygons() []polygon {
	var p []polygon
	p = append(p, cn.polygons...)
	if cn.front != nil {
		p = append(p, cn.front.allPolygons()...)
	}
	if cn.back != nil {
		p = append(p, cn.back.allPolygons()...)
	}
	return p
}

func (cn *csgNode) clipPolygons(polygons []polygon) []polygon {
	if cn.plane == nil {
		polygonsCopy := make([]polygon, len(polygons))
		copy(polygonsCopy, polygons)
		return polygonsCopy
	}
	var front, back []polygon
	for _, poly := range polygons {
		cn.plane.splitPolygon(poly, &front, &back, &front, &back)
	}
	if cn.front != nil {
		cn.front.clipPolygons(polygons)
	}
	if cn.back != nil {
		cn.back.clipPolygons(polygons)
	} else {
		back = []polygon{}
	}
	return append(front, back...)
}

func (cn *csgNode) clipTo(bsp *csgNode) {
	cn.polygons = bsp.clipPolygons(cn.polygons)
	if cn.front != nil {
		cn.front.clipTo(bsp)
	}
	if cn.back != nil {
		cn.back.clipTo(bsp)
	}
}

func (cn *csgNode) build(polygons []polygon) {
	if len(polygons) == 0 {
		return
	}
	if cn.plane == nil {
		cn.plane = &polygons[0].plane
	}
	var front, back []polygon
	for _, poly := range polygons {
		cn.plane.splitPolygon(poly, &cn.polygons, &cn.polygons, &front, &back)
	}
	if len(front) > 0 {
		if cn.front == nil {
			cn.front = &csgNode{}
		}
		cn.front.build(front)
	}
	if len(back) > 0 {
		if cn.back == nil {
			cn.back = &csgNode{}
		}
		cn.back.build(back)
	}
}
