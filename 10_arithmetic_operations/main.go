package main

import "fmt"

func main() {
	// Toplama işlemi:
	// + operatörü iki sayıyı toplar.
	var a int = 10
	var b int = 20

	fmt.Println("Toplama:", a+b) // 10 + 20 = 30

	// Çıkarma işlemi:
	// - operatörü bir sayıdan diğerini çıkarır.
	var x int = 30
	var y int = 10

	fmt.Println("Çıkarma:", x-y) // 30 - 10 = 20

	// Çarpma işlemi:
	// * operatörü iki sayıyı çarpar.
	var p int = 5
	var q int = 6

	fmt.Println("Çarpma:", p*q) // 5 * 6 = 30

	// Bölme işlemi:
	// / operatörü bir sayıyı diğerine böler.
	var m int = 10
	var n int = 4

	fmt.Println("Bölme:", m/n) // 10 / 4 = 2 (kesirli kısmı atılır)

	// Modülüs işlemi:
	// % operatörü iki sayının bölümünden kalan değeri verir.
	var u int = 10
	var v int = 3

	fmt.Println("Modülüs:", u%v) // 10 % 3 = 1 (kalan 1)

	// Artırma(Increment) işlemi:
	// ++ operatörü bir sayıyı 1 artırır.
	var increment int = 5
	increment++                        // increment değeri 1 artar.
	fmt.Println("Artırma:", increment) // 5 + 1 = 6

	// Azaltma(Decrement) işlemi:
	// -- operatörü bir sayıyı 1 azaltır.
	var decrement int = 5
	decrement-- // decrement değeri 1 azalır.
	// fmt.Println("Azaltma:" decrement++) Bu satır hata verir. Çünkü Go dilinde Increment ve Decrement birer statement dır. Bir satırda sadece bir statement bulunabilir.

	fmt.Println("Azaltma:", decrement) // 5 - 1 = 4
}
