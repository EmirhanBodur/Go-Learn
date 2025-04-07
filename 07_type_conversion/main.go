package main

import "fmt"

func main() {

	//var sayi1 = 10   // sayi1 bir int türünde (tam sayı)
	//var sayi2 = 20.0 // sayi2 bir float64 türünde (ondalıklı sayı)
	// sayi1 ve sayi2 farklı türlere sahip olduğundan, bunları doğrudan toplamak mümkün değildir.
	// Bu durumda Go, bir int ile float64 türünde toplama işlemi yapmaya çalıştığında hata verir.

	/* fmt.Println(sayi1 + sayi2) */
	// Ancak sayi2'yi int türüne yada sayi1'i float64 türüne dönüştürerek işlemi gerçekleştirebiliriz.
	// Dönüştürme işlemi (int(sayi2)), sayi2'nin ondalık kısmını kaybeder.
	// Örneğin 20.0 sayısı 20'ye dönüşür.
	// fmt.Println(sayi1 + int(sayi2)) // Çıktı: 30

	// int tipinde bir değişken
	var age int = 25

	// int → float64 dönüşümü
	var ageInFloat float64 = float64(age)

	// float64 → int dönüşümü (ondalık kısmı kaybolur!)
	var height float64 = 172.5
	var heightInt int = int(height)

	// uint → int dönüşümü (işaretli hale gelir)
	var score uint = 90
	var scoreInt int = int(score)

	// Ekrana yazdırma
	fmt.Println("Yaş (int):", age)
	fmt.Println("Yaş (float64):", ageInFloat)
	fmt.Println("Boy (float64):", height)
	fmt.Println("Boy (int):", heightInt)
	fmt.Println("Puan (uint):", score)
	fmt.Println("Puan (int):", scoreInt)
}
