package csg

import (
	"math"
)

type vector3 struct {
	X, Y, Z float64
}

func (v vector3) negated() vector3 {
	return vector3{-v.X, -v.Y, -v.Z}
}

func (v vector3) plus(a vector3) vector3 {
	return vector3{v.X + a.X, v.Y + a.Y, v.Z + a.Z}
}

func (v vector3) minus(a vector3) vector3 {
	return vector3{v.X - a.X, v.Y - a.Y, v.Z - a.Z}
}

func (v vector3) times(a float64) vector3 {
	return vector3{v.X * a, v.Y * a, v.Z * a}
}

func (v vector3) dividedBy(a float64) vector3 {
	return vector3{v.X / a, v.Y / a, v.Z / a}
}

func (v vector3) dot(a vector3) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
}

func (v vector3) lerp(a vector3, t float64) vector3 {
	return v.plus(a.minus(v).times(t))
}

func (v vector3) length() float64 {
	return math.Sqrt(v.dot(v))
}

func (v vector3) unit() vector3 {
	return v.dividedBy(v.length())
}

func (v vector3) cross(a vector3) vector3 {
	return vector3{
		X: v.Y*a.Z - v.Z*a.Y,
		Y: v.Z*a.X - v.X*a.Z,
		Z: v.X*a.Y - v.Y*a.X,
	}
}
