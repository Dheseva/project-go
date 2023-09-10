package main

import "fmt"

func main() {

	var result []string
	var arrayOne []string
	var arrayTwo []string
	arrayCheck := false
	fmt.Println("Penggabungan Array")
	// fmt.Println("Array Pertama")
	// fmt.Println("Array Kedua")

	var n1 int
	var n2 int

	fmt.Print("Masukkan jumlah elemen dalam array pertama: ")
	fmt.Scan(&n1)

	// Buat slice untuk menyimpan elemen-elemen string
	arrayOne = make([]string, n1)

	// Masukkan input dari pengguna ke dalam slice
	for i := 0; i < n1; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arrayOne[i])
	}

	fmt.Print("Masukkan jumlah elemen dalam array kedua: ")
	fmt.Scan(&n2)

	// Buat slice untuk menyimpan elemen-elemen string
	arrayTwo = make([]string, n2)

	// Masukkan input dari pengguna ke dalam slice
	for i := 0; i < n2; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arrayTwo[i])
	}

	result, arrayCheck = arrayMerge(arrayOne[:],arrayTwo[:])
	// result, arrayCheck = arrayMerge([]string{}, []string{"rusia","jepang","devil"})

	if !arrayCheck{
		fmt.Println("Nama Array ada yg sama")
	}else{
		fmt.Println(result)
	}

	fmt.Println("Menghitung array pada slice")
	frequency := Mapping([]string{"ass","asd","a", "b", "c", "a", "b", "d", "a", "e"})
	fmt.Println("Frekuensi kemunculan setiap nilai:")
	for key, count := range frequency {
		fmt.Printf("%s: %d\n", key, count)
	}

}

func arrayMerge(array1, array2 []string) ([]string, bool) {

	var arMerge []string
	for i:= 0 ; i < len(array1); i++{
		for v:= 0 ; v < len(array2); v++{
			if array1[i] == array2[v]{
				return []string{"_"} ,false
			}
		}
	}

	arMerge = append(array1, array2...)
	return arMerge, true
}

func Mapping(slices []string) map[string]int {
	allKeys := make(map[string]int)
	// list := []string{}
	for _, values := range slices {
            allKeys[values] ++
    }
    return allKeys
}