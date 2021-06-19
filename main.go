package main

import "fmt"

func main() {

	//Best Order Algorithm: Quick Sort
	slice := []int8{8, 7, 9, 4, 5, 6, 3, 1, 10, 2}

	fmt.Printf("Debbug1 ~Before to QuickSort:\n~~~%v~~~\n\n", slice)

	quicksort(slice, 0, int8(len(slice)-1))

	fmt.Printf("Debbug2 ~After to QuickSort:\n~~~%v~~~\n\n", slice)

	//Best Search Algorithm: Binary Search
	indexObjetive := binarySearch(slice, 0, int8(len(slice)-1), 6)

	fmt.Printf("Debbug3 ~After to BinarySearch:\n~~~Index:%d - Value:%d~~~\n\n", indexObjetive, slice[indexObjetive])
}

func quicksort(slice []int8, lo, hi int8) {
	if lo > hi {
		return
	}

	indexPivot := partition(slice, lo, hi)

	quicksort(slice, lo, indexPivot-1)
	quicksort(slice, indexPivot+1, hi)

}

func partition(slice []int8, lo, hi int8) int8 {
	pivot := slice[hi]

	for i := lo; i < hi; i++ {
		if slice[i] < pivot {
			slice[i], slice[lo] = slice[lo], slice[i]
			lo++
		}
	}

	slice[lo], slice[hi] = slice[hi], slice[lo]

	return lo
}

func binarySearch(slice []int8, left, right, objetive int8) int8 {

	for left <= right {
		var mid = left + (right-left)/2

		if slice[mid] == objetive {
			return mid
		}

		if slice[mid] < objetive {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
