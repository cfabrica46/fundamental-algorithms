package main

import "testing"

var sliceBase = []int8{8, 7, 9, 4, 5, 6, 3, 1, 10, 2}
var repetitions = 10000000

func TestQuickSortRecursive(t *testing.T) {
	var slice []int8
	for i := 0; i < repetitions; i++ {
		slice = sliceBase
		QuicksortRecursive(slice, 0, int8(len(slice)-1))
	}
}

func TestQuickSortIteraive(t *testing.T) {
	var slice []int8
	for i := 0; i < repetitions; i++ {
		slice = sliceBase
		QuicksortIterative(slice, 0, int8(len(slice)-1))
	}
}
