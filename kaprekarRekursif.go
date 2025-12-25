package main

import (
	"fmt"
)

func dsc(num int) int {
	d1 := num / 1000
	d2 := (num / 100) % 10
	d3 := (num / 10) % 10
	d4 := num % 10

	if d1 < d2 { d1, d2 = d2, d1 }
	if d1 < d3 { d1, d3 = d3, d1 }
	if d1 < d4 { d1, d4 = d4, d1 }

	if d2 < d3 { d2, d3 = d3, d2 }
	if d2 < d4 { d2, d4 = d4, d2 }

	if d3 < d4 { d3, d4 = d4, d3 }

	return d1*1000 + d2*100 + d3*10 + d4
}

func asc(num int) int {
	d1 := num / 1000
	d2 := (num / 100) % 10
	d3 := (num / 10) % 10
	d4 := num % 10

	if d1 > d2 { d1, d2 = d2, d1 }
	if d1 > d3 { d1, d3 = d3, d1 }
	if d1 > d4 { d1, d4 = d4, d1 }

	if d2 > d3 { d2, d3 = d3, d2 }
	if d2 > d4 { d2, d4 = d4, d2 }

	if d3 > d4 { d3, d4 = d4, d3 }

	return d1*1000 + d2*100 + d3*10 + d4
}

func validasiInput(num int) bool {
	if num < 0 || num > 9999 {
		return false
	}

	d1 := num / 1000
	d2 := (num / 100) % 10
	d3 := (num / 10) % 10
	d4 := num % 10

	if d1 == d2 && d2 == d3 && d3 == d4 {
		return false
	}

	return true
}

func kaprekarRecursive(currentNum int, step int) int {
	if currentNum == 6174 {
		return step
	}

	ascNum := asc(currentNum)
	dscNum := dsc(currentNum)
	result := dscNum - ascNum

	if ascNum < 1000 {
		fmt.Printf("Langkah %d : 0%d - %d = %d\n", step+1, ascNum, dscNum, result)
	} else {
		fmt.Printf("Langkah %d : %d - %d = %d\n", step+1, ascNum, dscNum, result)
	}

	return kaprekarRecursive(result, step+1)
}

func main() {
	var input int

	fmt.Print("Masukkan: ")
	fmt.Scan(&input)
	fmt.Println("==========================================")
	if !validasiInput(input) {
		fmt.Println("Error: Input tidak valid!")
		fmt.Println("Pastikan:")
		fmt.Println("1. Angka antara 0 - 9999")
		fmt.Println("2. Minimal memiliki 2 digit berbeda")
		fmt.Println("3. Bukan angka seperti 1111, 2222, dll.")
		return
	}
	fmt.Println("Memproses angka:", input)
	fmt.Println("==========================================")

	fmt.Println("Ascending: ", asc(input))
	fmt.Println("Descending:", dsc(input))
	fmt.Println("==========================================")

	if input == 6174 {
		fmt.Println("Angka sudah merupakan konstanta Kaprekar (6174)")
		fmt.Println("==========================================")
		fmt.Println("Dibutuhkan 0 iterasi untuk mencapai 6174")
	} else {
		fmt.Println("Langkah-langkah perhitungan:")
		iterasi := kaprekarRecursive(input, 0)
		fmt.Println("==========================================")
		fmt.Printf("Dibutuhkan %d iterasi untuk mencapai 6174\n", iterasi)
	}
}
