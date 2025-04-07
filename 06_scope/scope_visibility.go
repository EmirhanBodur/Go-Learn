package main

import "fmt"

var globalVar = "Global Değişken"

// globalVar := "Global Değişken" ❌ Bu kısa yol(:=) ile belirttiğim değişkeni fonksiyon içinde kullanamayız çünkü Go bunun türünü belirleyemez ve hata mesajı verir.

// Şimdi akla şu soru gelebilir:
// Neden tüm değişkenleri fonksiyon dışında global olarak tanımlamıyoruz?
// Çünkü global değişkenler uygulama süresince bellekte kalır,
// ancak local değişkenler sadece fonksiyon çalışırken var olur.
// Bu da bellek kullanımı ve performans açısından daha verimlidir.

func main() { // Burada main'in altını çizme sebebi klasör içindeki oluşturduğumuz dosyalarda birden fazla main fonksiyonu kullandığımız için altını kırmızı çiziyor.
	// Bu konuları öğrenme aşaması olduğu için hepsine main
	// fonksiyonu oluşturdum fakat gerçek projelerde sadece ana bir main fonksiyonu kullanılmalı ve diğerlerine farklı bir fonksiyon adı verilmelidir.
	localVar := "Local Değişken"
	fmt.Println(globalVar) // ✅ Global erişilir
	fmt.Println(localVar)  // ✅ Local erişilir

	if true {
		blockVar := "Block Değişken"
		fmt.Println(blockVar) // ✅ Bu blokta erişilir
	}
	// fmt.Println(blockVar) // ❌ Dışarıda kullanılamaz

}
