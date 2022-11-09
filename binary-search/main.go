package binarysearch

func BinarySearch(slice []int8, left, right, objetive int8) int8 {
	for left <= right {
		mid := left + (right-left)/2

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
