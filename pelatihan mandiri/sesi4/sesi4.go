package main

import (
	"fmt"
	"slices"
)

type Car struct {
	Type   string
	FuelIn float32
}

type Student struct{
	Name string
	Nilai int
}

func (cars Car) Reciver() {

	Mil := cars.FuelIn / 1.5

	fmt.Printf("Car %s telah menempuh %f Mil\n", cars.Type, Mil)
}

func (siswa Student) Memamix() {



	Scores := []int{siswa.Nilai}
	Names := []string{siswa.Name}
	Max := slices.Max(Scores)
	Min := slices.Min(Scores)
	var MaxSiswa string
	var MinSiswa string

	for k:= 0; k < len(Scores);k++{
		if Max == siswa.Nilai {
			MaxSiswa = Names[k]
		}else if Min == siswa.Nilai {
			MinSiswa = Names[k]
		}
	}
	sum := 0
	for _, score := range Scores {
		sum += score
	}
	average := float64(sum) / float64(len(Scores))

	//fmt.Println(len(Scores))
	fmt.Printf("Average Score %f\n",average)
	fmt.Printf("Max Score %d is %s\n", Max, MaxSiswa)
	fmt.Printf("Min Score %d is %s\n", Min, MinSiswa)
	// fmt.Println()
}

func main() {

	// Method Reciver

	var t string
	var l float32

	fmt.Printf("Type Cars = ")
	fmt.Scanln(&t)
	fmt.Printf("Jumlah fuel Cars (L) = ")
	fmt.Scanln(&l)

	cars := Car{
		Type: t,
		FuelIn: l,
	}

	cars.Reciver()

	//// median , max , min
	// var s string
	// var n int
	var js int
	var siswaa []Student

	fmt.Printf("Jumlah student = ")
	fmt.Scanln(&js)

	siswas := make([]Student, js)

	if js > 0 {
		for p:= 0;p < js;p++{
			var ss Student
		fmt.Printf("Nama student %d = ",p+1)
		fmt.Scanln(&ss.Name)
		fmt.Printf("Jumlah score %d = ",p+1)
		fmt.Scanln(&ss.Nilai)

		//siswas[p] = ss

		siswaa[p] = Student{
			Name: ss.Name,
			Nilai:  ss.Nilai,
			}
		}
		
		// siswaa := []Student{

		// }
		fmt.Println(siswas)
		siswaa.Memamix()
	}

	
}