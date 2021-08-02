package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e slices")

	var array [5]string
	array[0] = "Posição 1"

	fmt.Println(array)

	array1 := [5]string{"po1", "po2", "po3", "po4", "po5"}

	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4, 5} //fixa o tamanho do array de acordo com os intens passados
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}

	fmt.Println(slice)

	slice = append(slice, 15)

	fmt.Println(slice)

	slice2 := array2[1:3] //serve como um ponteiro

	fmt.Println(slice2)

	/*fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))*/

	//Arrays internos
	slice3 := make([]float64, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 1, 2, 3, 3, 4, 5, 6, 7)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float64, 5)

	fmt.Println(slice4)
	slice4 = append(slice4, 3)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
