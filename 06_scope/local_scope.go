package main

import "fmt"

func main() { // Burada main'in altını çizme sebebi klasör içindeki oluşturduğumuz dosyalarda birden fazla main fonksiyonu kullandığımız için altını kırmızı çiziyor.
	// Bu konuları öğrenme aşaması olduğu için hepsine main
	// fonksiyonu oluşturdum fakat gerçek projelerde sadece ana bir main fonksiyonu kullanılmalı ve diğerlerine farklı bir fonksiyon adı verilmelidir.
	age := 22                // Local değişken, sadece main fonksiyonunda geçerlidir
	fmt.Println("Yaş:", age) // Local değişken sadece bu fonksiyonda geçerli
}

func showAge() {
	// fmt.Println(age) // ❌ Hata verir, çünkü 'age' sadece main fonksiyonu içinde tanımlanmıştır
}
