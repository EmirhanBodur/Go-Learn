package main

import "fmt"

// Typeless Constants, Go dilinde tür belirtmeden sabitler tanımlamanıza olanak tanır.
// Go dilinde sabitlerin türü, sabitin değerine göre otomatik olarak belirlenir.
// Bu, genellikle sabitlerin türe özgü tanımlama gereksinimlerini ortadan kaldırır ve
// Go'nun tip çıkarımını kullanır.

func main() {

	const x = 3
	var y int16 = 12

	fmt.Printf("%T, %v\n", x, x) // x'in türü int, değeri 3
	fmt.Printf("%T, %v\n", y, y) // y'nin türü int16, değeri 12

	fmt.Printf("%T, %v\n", x+y, x+y) // Go bu işlemi otomatik olarak int16 türünde tanımlar

	// Tekrar x'in türünü ve değerini yazdıralım
	fmt.Printf("%T, %v\n", x, x) // x'in türü int, değeri 3
}
