package csg

const planeEpsilon = 1e-5

type plane struct {
	normal vector3
	w      float64
}

func planeFromPoints(a, b, c vector3) plane {
	n := b.minus(a).cross(c.minus(a)).unit()
	return plane{n, n.dot(a)}
}

func (p plane) flipped() plane {
	return plane{
		p.normal.negated(),
		-p.w,
	}
}

/*
splitPolygon splits the polygon into lists which are returned in the same order as the arguments.
*/
func (p plane) splitPolygon(poly polygon, coplanarFront, coplanarBack, front, back *[]polygon) {
	const (
		coplanarType = 0
		frontType    = 1
		backType     = 2
		spanningType = 3
	)
	polygonType := 0
	var types []int
	for _, v := range poly.vertices {
		t := p.normal.dot(v.pos) - p.w
		var pType int
		if t < -planeEpsilon {
			pType = backType
		} else {
			if t > planeEpsilon {
				pType = frontType
			} else {
				pType = coplanarType
			}
		}
		polygonType |= pType
		types = append(types, pType)
	}
	switch polygonType {
	case coplanarType:
		if p.normal.dot(poly.plane.normal) > 0 {
			*coplanarFront = append(*coplanarFront, poly)
		} else {
			*coplanarBack = append(*coplanarBack, poly)
		}
	case frontType:
		*front = append(*front, poly)
	case backType:
		*back = append(*back, poly)
	case spanningType:
		var f, b []vertex
		for i, vi := range poly.vertices {
			j := (i + 1) % len(poly.vertices) // next vertex of polygon (wraps over)
			ti := types[i]
			tj := types[j]
			vj := poly.vertices[j]
			if ti != backType {
				f = append(f, vi)
			}
			if ti != frontType {
				b = append(b, vi)
			}
			if (ti | tj) == spanningType {
				t := (p.w - p.normal.dot(vi.pos)) / p.normal.dot(vj.pos.minus(vi.pos))
				v := vi.interpolated(vj, t)
				f = append(f, v)
				b = append(b, v)
			}
		}
		if len(f) >= 3 {
			*front = append(*front, newPolygon(f, poly.shared))
		}
		if len(b) >= 3 {
			*back = append(*back, newPolygon(b, poly.shared))
		}
	}
}
