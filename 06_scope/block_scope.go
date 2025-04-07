package main

import "fmt"

func main() { // Burada main'in altını çizme sebebi klasör içindeki oluşturduğumuz dosyalarda birden fazla main fonksiyonu kullandığımız için altını kırmızı çiziyor.
	// Bu konuları öğrenme aşaması olduğu için hepsine main
	// fonksiyonu oluşturdum fakat gerçek projelerde sadece ana bir main fonksiyonu kullanılmalı ve diğerlerine farklı bir fonksiyon adı verilmelidir.
	if true {
		message := "Merhaba"
		fmt.Println(message) // ✅ Bu blok içinde geçerli
	}
	// fmt.Println(message) // ❌ Blok dışında kullanılamaz
}
