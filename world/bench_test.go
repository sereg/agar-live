package world

//func BenchmarkStruct(b *testing.B) {
//	grid := newGrid(50)
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		grid = newGrid(50)
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.set(i * 40, j * 40, int(i))
//			}
//		}
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.getObjInVision(i * 40, j * 40, 70)
//			}
//		}
//	}
//	//fmt.Printf("Struct count - %d", len(grid.data))
//}
//
//func BenchmarkString(b *testing.B) {
//	grid := newGrid(50)
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		grid = newGrid(50)
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.setString(i * 40, j * 40, int(i))
//			}
//		}
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.getObjInVisionString(i * 40, j * 40, 70)
//			}
//		}
//	}
//	//fmt.Printf("String count - %d", len(grid.dataString))
//}
//
//func BenchmarkMultiply(b *testing.B) {
//	grid := newGrid(50)
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		grid = newGrid(50)
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.setMultiply(i * 40, j * 40, int(i))
//			}
//		}
//		for i:= 0.0; i < 100; i++ {
//			for j := 0.0; j < 100; j++ {
//				grid.getObjInVisionMultiply(i * 40, j * 40, 70)
//			}
//		}
//	}
//	//fmt.Printf("Multiply count - %d", len(grid.dataMultiply))
//}