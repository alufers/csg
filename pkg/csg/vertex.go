package csg

type vertex struct {
	pos    vector3
	normal vector3
}

func (v vertex) flipped() vertex {
	return vertex{
		pos:    v.pos,
		normal: v.normal.negated(),
	}
}

func (v vertex) interpolated(other vertex, t float64) vertex {
	return vertex{
		pos:    v.pos.lerp(other.pos, t),
		normal: v.normal.lerp(other.normal, t),
	}
}
