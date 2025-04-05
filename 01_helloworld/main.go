// Uygulamanın ana paketi
package main

// fmt: Formatlı giriş/çıkış işlemleri için kullanılan standart kütüphane
import "fmt"

// Programın başlangıç fonksiyonu
func main() {
	var name string // Kullanıcının girişi için bir string değişken tanımlandı

	fmt.Println("Selam GitHub Ben Emirhan Bodur") // Ekrana düz yazı yazar
	fmt.Printf("Selam, %s!\n", "Emirhan")         // Formatlı yazı yazar, %s yerine "Emirhan" gelir
	fmt.Scan(&name)                               // Kullanıcıdan veri alır ve name değişkenine atar
	fmt.Println("Merhaba", name)                  // Kullanıcının adını ekrana yazdırır
}
