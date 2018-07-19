package csg

// func TestClipPolygons(t *testing.T) {
// 	sp1node := newCsgNode(makeSphere(vector3{0, 0, 0}, 2).polygons)
// 	polys := sp1node.clipPolygons([]polygon{
// 		polygon{
// 			vertices: []vertex{
// 				{
// 					pos: vector3{
// 						0.7653668647301796,
// 						1.8477590650225735,
// 						0,
// 					},
// 					normal: vector3{
// 						-0.3826834323650898,
// 						-0.9238795325112867,
// 						0,
// 					},
// 				},
// 				{
// 					pos: vector3{
// 						0.7071067811865476,
// 						1.8477590650225735,
// 						0.2928932188134525,
// 					},
// 					normal: vector3{
// 						-0.3535533905932738,
// 						-0.9238795325112867,
// 						-0.14644660940672624,
// 					},
// 				},
// 				{
// 					pos:    vector3{0, 2, 0},
// 					normal: vector3{0, -1, 0},
// 				},
// 			},
// 			plane: plane{
// 				normal: vector3{
// 					-0.19494359482964702,
// 					-0.9800476330240645,
// 					-0.0387766919530659,
// 				},
// 				w: -1.960095266048129,
// 			},
// 		},
// 	})
// 	t.Logf("%+v", polys)
// 	t.Fail()
// }
