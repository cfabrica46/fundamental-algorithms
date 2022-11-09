package recursive

func QuicksortRecursive(slice []int8, lo, hi int8) {
	if lo > hi {
		return
	}

	indexPivot := partitionRecursive(slice, lo, hi)

	QuicksortRecursive(slice, lo, indexPivot-1)
	QuicksortRecursive(slice, indexPivot+1, hi)
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
