package main

import (
	"fmt"

	"github.com/cfabrica46/fundamental-algorithms/recursive"
)

func main() {
	arr := []int8{4, 1, 2, 19, 40, 21}

	recursive.Quicksort(arr, 0, int8(len(arr))-1)

	fmt.Println(arr)
}
