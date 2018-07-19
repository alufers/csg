package csg

import (
	"os"
)

// Run is the real entrypoint
func Run() {
	sp1 := makeSphere(vector3{0, 0, 0}, 2)
	sp2 := makeSphere(vector3{0, 0, 3}, 2)
	//fmt.Println(len(sp1.subtract(sp2).polygons))
	sp1.subtract(sp2).toObj(os.Stdout)

}
