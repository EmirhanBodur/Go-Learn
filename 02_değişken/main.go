package main

import "fmt"

func main() {
	// name adında bir string değişken tanımlanır ve "Emirhan" değeri atanır
	var ad string = "Emirhan"

	// Tür belirtilmeden tanımlanır, Go otomatik olarak int atar
	var yas = 22

	// Kısa yol (Sadece fonksiyon içinde kullanılır)
	soyad := "Bodur"

	var name, surname string = "Emirhan", "Bodur"

	
	fmt.Println(ad)
	fmt.Println(soyad)
	fmt.Println(yas)
	fmt.Println("Ad:", name ,"Soyad:", surname)
}
