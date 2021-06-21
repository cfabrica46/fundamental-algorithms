package main

import "fmt"

func main() {

	//Best Order Algorithm: Quick Sort
	slice := []int8{8, 7, 9, 4, 5, 6, 3, 1, 10, 2}

	fmt.Printf("Debbug1 ~Before to QuickSort:\n~~~%v~~~\n\n", slice)

	//quicksortRecursive(slice, 0, int8(len(slice)-1))

	quicksortIterative(slice, 0, int8(len(slice)-1))

	fmt.Printf("Debbug2 ~After to QuickSort:\n~~~%v~~~\n\n", slice)

	//Best Search Algorithm: Binary Search
	indexObjetive := binarySearch(slice, 0, int8(len(slice)-1), 6)

	fmt.Printf("Debbug3 ~After to BinarySearch:\n~~~Index:%d - Value:%d~~~\n\n", indexObjetive, slice[indexObjetive])
}

func quicksortIterative(slice []int8, lo, hi int8) {

	stack := make([]int8, hi-lo+1)

	var top = -1

	top++
	stack[top] = lo
	top++
	stack[top] = hi

	for top >= 0 {
		hi = stack[top]
		top--
		lo = stack[top]
		top--

		p := partitionIterative(slice, lo, hi)

		if p-1 > lo {
			top++
			stack[top] = lo
			top++
			stack[top] = p - 1
		}

		if p+1 < hi {
			top++
			stack[top] = p + 1
			top++
			stack[top] = hi
		}
	}

}

func partitionIterative(slice []int8, lo, hi int8) int8 {

	var pivot = slice[hi]

	var aux = lo - 1

	for i := lo; i <= hi-1; i++ {
		if slice[i] <= pivot {
			aux++

			slice[i], slice[aux] = slice[aux], slice[i]
		}
	}

	slice[aux+1], slice[hi] = slice[hi], slice[aux+1]

	return aux + 1
}

func quicksortRecursive(slice []int8, lo, hi int8) {
	if lo > hi {
		return
	}

	indexPivot := partitionRecursive(slice, lo, hi)

	quicksortRecursive(slice, lo, indexPivot-1)
	quicksortRecursive(slice, indexPivot+1, hi)

}

func partitionRecursive(slice []int8, lo, hi int8) int8 {
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
