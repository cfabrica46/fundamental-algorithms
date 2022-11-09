package iterative

func QuicksortIterative(slice []int8, lo, hi int8) {
	stack := make([]int8, hi-lo+1)

	top := -1

	top++

	stack[(top)] = lo

	top++

	stack[top] = hi

	for top >= 0 {
		hi = stack[top]
		top--
		lo = stack[top]
		top--

		partition := partitionIterative(slice, lo, hi)

		if partition-1 > lo {
			top++

			stack[top] = lo

			top++

			stack[top] = partition - 1
		}

		if partition+1 < hi {
			top++

			stack[top] = partition + 1

			top++

			stack[top] = hi
		}
	}
}

func partitionIterative(slice []int8, lo, hi int8) int8 {
	pivot := slice[hi]

	aux := lo - 1

	for i := lo; i <= hi-1; i++ {
		if slice[i] <= pivot {
			aux++

			slice[i], slice[aux] = slice[aux], slice[i]
		}
	}

	slice[aux+1], slice[hi] = slice[hi], slice[aux+1]

	return aux + 1
}
