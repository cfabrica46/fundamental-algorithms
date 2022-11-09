package main

//func BenchmarkQuickSort(b *testing.B) {
//	b.ResetTimer()
//	var slice []int8
//	for i := 0; i < b.N; i++ {
//		slice = sliceBase
//		QuicksortRecursive(slice, 0, int8(len(slice)-1))
//	}
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		slice = sliceBase
//		QuicksortIterative(slice, 0, int8(len(slice)-1))
//	}
//	b.ReportAllocs()
//
//}

//func BenchmarkQuickSortIterative(b *testing.B) {
//	b.ResetTimer()
//	var slice []int8
//	for i := 0; i < b.N; i++ {
//		slice = sliceBase
//		QuicksortIterative(slice, 0, int8(len(slice)-1))
//	}
//}

// go test -run=QuickSortRecursive -bench=.
// go test -run=QuickSortIterative -bench=.
