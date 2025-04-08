package main

import "fmt"

func main() {
	// CamelCase ile ilgili ipuçları:
	// - Değişken ve fonksiyon isimlerini kısa ve öz tutmalıyız. Çünkü bir ekip çalışması yaptığımızda başkası kodunuzu incelerken değişken ve fonksiyon ismine baktığında ne amaçla yazıldığınız anlaması gerekiyor.
	//   Anlamlı ve işlevsel olmasına dikkat edin.
	//   Örnek: 'getUserDetails' doğru bir isimdir, 'getAllUserDetailsFromDatabase' gereksiz uzun bir isimdir.
	//
	// - Rastgele harfler kullanmaktan kaçının. 'a', 'b', 'c' gibi isimler yerine anlamlı ve açıklayıcı isimler tercih edin.
	//   Örnek: 'userName', 'totalAmount', 'calculateSalary' gibi.
	//
	// - Birden fazla kelimeyi birleştirirken her kelimenin anlamlı olduğundan emin olun.
	//   Örnek: 'fetchUserInfoFromAPI' doğru bir kullanımken, 'getUserInfoFromWebAPI' gereksiz uzunluktadır.
	//
	// CamelCase Kuralları:
	// - Değişkenler ve fonksiyonlar için küçük harfle başlar (lowerCamelCase): 'userName', 'calculateSum'
	// - Yapılar, tipler, dosya isimleri ve public öğeler için büyük harfle başlar (UpperCamelCase): 'ProductList', 'EmployeeDetails'

	var userName string = "Emirhan" // 'userName' küçük harf ile başlayarak CamelCase'e uygun bir isimlendirme
	var userAge int = 22            // 'userAge' anlamlı ve kısa bir isimlendirme
	var isStudent bool = true       // 'isStudent' doğru bir CamelCase kullanımı
	var averageScore float32 = 88.5 // 'averageScore' yine uygun bir isimlendirme

	// fmt.Print: Yeni satıra geçmez.
	fmt.Print("İsim: ")
	fmt.Print(userName)

	// fmt.Println: Yeni satıra geçer
	fmt.Println("\nYaş:", userAge)

	// fmt.Printf: Formatlı yazım
	// Formatlı yazımda % işaretiyle tür belirtilir:
	// %s: String formatı, metin türündeki veriler için
	// %d: Integer formatı, tam sayı verileri için
	// %f: Float formatı, ondalıklı sayılar için
	// %t: Boolean formatı, doğru/yanlış değerleri için
	// %v: Varsayılan format, veri türüne göre uygun formatta yazdırır
	// %+v: Yapı (struct) verisini içerik olarak gösterir
	// %#v: Yapının go kodu şeklinde gösterimi
	// %p: Pointer türündeki bir değişkenin adresini gösterir
	fmt.Printf("Öğrenci misin? %t\n", isStudent)
	fmt.Printf("Ortalama: %f\n", averageScore)
}
