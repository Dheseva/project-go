package main

import "fmt"

func main() {

	bp := 0
	isPrima := false

	fmt.Println("Menentukan Bilangan Prima")
	fmt.Println("Masukkan Bilangan : ")
	fmt.Scanln(&bp)
	fmt.Println("Bilangan ",bp)
	if bp != 0 {
		if bp < 1 || bp == 1 || bp == 2 {
			isPrima = true
		}else if bp % 2 == 0{

		}else{
			isPrima = bilPrima(bp)
		}
		
	}

	if isPrima {
		fmt.Println("Bilangan Prima")
	}else{
		fmt.Println("Bilangan bukan Prima")
	}

	var k7 int
	isKelipatan := false
	fmt.Println("Menentukan Bilangan Kelipatan 7")
	fmt.Println("Masukkan Bilangan : ")
	fmt.Scanln(&k7)

	if k7 > 0 {
		if k7 % 7 == 0{
			isKelipatan = true
		}
	}

	if isKelipatan {
		fmt.Println("Bilangan Kelipatan")
	}else{
		fmt.Println("Bilangan bukan Kelipatan")
	}
}

func bilPrima(bp int) bool {
	
	for i:= 3; i < bp;i += 2{
		if bp % i == 0 {

			isPrima := false
			return isPrima
		}
	}
	isPrima := true
	return isPrima
}