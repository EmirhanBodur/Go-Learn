package main

import "fmt"

func main() {
	// Karşılaştırma Operatörleri (Comparison Operators)
	// Bu operatörler, iki değeri karşılaştırır ve boolean (`true` veya `false`) bir değer döndürür.

	// Eşit mi? (==)
	a := 10
	b := 20
	fmt.Println("Eşit mi?:", a == b) // false

	// Eşit değil mi? (!=)
	fmt.Println("Eşit değil mi?:", a != b) // true

	// Büyüktür (>)
	fmt.Println("Büyük mü?:", a > b) // false

	// Küçüktür (<)
	fmt.Println("Küçük mü?:", a < b) // true

	// Büyük veya eşit mi? (>=)
	fmt.Println("Büyük veya eşit mi?:", a >= b) // false

	// Küçük veya eşit mi? (<=)
	fmt.Println("Küçük veya eşit mi?", a <= b) // true
	loglicalOperators()
}

func loglicalOperators() {
	// Mantıksal Operatörler (Logical Operators)
	// Bu operatörler, bir veya daha fazla koşulu birleştirmek için kullanılır.

	a := true
	b := false
	// Bu operatör, her iki durumun da true olması halinde true döner.
	// VE (AND) (&&) Operatörü
	// 'a' true ve 'b' false olduğu için sonuç false olacaktır.
	fmt.Println(a && b) // false, çünkü b false

	// VEYA (OR) (||) Operatörü
	// Bu operatör, iki durumdan birinin true olması durumunda true döner.
	// Yani, her iki koşuldan biri true ise sonuç true olacaktır.
	fmt.Println(a || b) // true, çünkü a true

	// DEĞİL (NOT) (!) Operatörü
	// Bu operatör, bir durumu tersine çevirir.
	// Eğer durum true ise false, false ise true yapar.
	fmt.Println(!a) // false, çünkü a true
}
