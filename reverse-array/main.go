package main

import "fmt"

func main() {
	arr := []int32{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr)
	fmt.Println(reverseArray(arr))
}

// Se requiere de un arreglo ordenado

// reverseArray repite el ciclo n/2 veces en donde n es la longitud
// del arreglo original

// [1,2,3,4,5,6,7] => [7,2,3,4,5,6,1] => [7,6,3,4,5,2,1] => [7,6,5,4,3,2,1]
func reverseArray(a []int32) []int32 {
	// i toma el valor del valor en la mitad menor del arreglo original
	for i := len(a)/2 - 1; i >= 0; i-- {
		// opp toma el valor de el opuesto del i esto debido a la siguinte propiedad
		// indiceMayor - indiceN = indiceOpuesto de N
		opp := len(a) - 1 - i

		// se reemplazan estos dos valores empezando desde la mitad
		// llegando a los extremos
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
