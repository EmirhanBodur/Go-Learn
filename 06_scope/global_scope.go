package main

import "fmt"

// Global değişken, tüm fonksiyonlar tarafından erişilebilir
var name = "Emirhan"

func main() { // Burada main'in altını çizme sebebi klasör içindeki oluşturduğumuz dosyalarda birden fazla main fonksiyonu kullandığımız için altını kırmızı çiziyor.
	// Bu konuları öğrenme aşaması olduğu için hepsine main
	// fonksiyonu oluşturdum fakat gerçek projelerde sadece ana bir main fonksiyonu kullanılmalı ve diğerlerine farklı bir fonksiyon adı verilmelidir.

	// Global değişken main fonksiyonunda kullanılıyor
	fmt.Println("Ad:", name) // Global değişken her yerden erişilebilir
	globalVariables()

}

func globalVariables() {
	fmt.Println("Soyad", name) // gördüğünüz gibi global değişkeni başka fonksiyonlarda da kullanabiliyoruz.
}
