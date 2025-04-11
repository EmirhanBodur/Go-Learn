package main

import "fmt"

func main() {
	// ----------------------
	// Basit if Yapısı
	// ----------------------

	x := 5
	if x > 0 {
		fmt.Println("X Büyüktür") // X pozitif olduğu için bu yazdırılır.
	}

	// ----------------------
	// If-Else Yapısı
	// ----------------------

	y := -3
	if y > 0 {
		fmt.Println("y pozitif bir sayıdır") // Bu yazdırılmaz çünkü y negatif.
	} else {
		fmt.Println("y negatif bir sayıdır.") // y negatif olduğu için burası yazdırılır.
	}

	// ----------------------
	// Else-If Yapısı
	// ----------------------

	z := 0
	if z > 0 {
		fmt.Println("Z pozitif bir sayıdır") // Z sıfır olduğu için burası yazdırılmaz.
	} else if z < 0 {
		fmt.Println("Z negatif bir sayıdır") // Z sıfır olduğu için burası yazdırılmaz.
	} else {
		fmt.Println("Z sıfırdır") // Z sıfır olduğu için bu yazdırılır.
	}

	// ----------------------
	// Örnek: Askere Gitme Durumu
	// ----------------------

	asker := 18
	var age int
	// Yaş bilgisini kullanıcıdan alıyoruz.
	fmt.Print("Yaşınızı giriniz: ")
	fmt.Scanln(&age)

	var gender string
	// Cinsiyet bilgisini kullanıcıdan alıyoruz.
	fmt.Print("Kadın mısınız, erkek misiniz? (Kadın/Erkek): ")
	fmt.Scanln(&gender)

	// Yaş 18 veya daha büyükse, erkekler askere gidebilir.
	if age >= asker {
		if gender == "Erkek" || gender == "erkek" {
			fmt.Println("Askere gidebilir")
		} else if gender == "Kadın" || gender == "kadın" {
			fmt.Println("Kadınlar askere gidemez")
		} else {
			fmt.Println("Cinsiyet bilgisi yanlış girildi. Tekrar Deneyiniz.")
		}
	} else {
		fmt.Println("Askere gidemez.") // Yaş 18'den küçükse askere gidilemez.
	}

	// ----------------------
	// Örnek 2: Artık Yıl Hesaplama
	// ----------------------

	var year int
	// Kullanıcıdan bir yıl alıyoruz.
	fmt.Print("Bir yıl giriniz: ")
	fmt.Scanln(&year)

	// Artık yıl olup olmadığını kontrol ediyoruz.
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "artık yıldır") // Eğer yıl 4 ile bölünebilir ama 100 ile bölünemez veya 400 ile bölünebiliyorsa, artık yıl.
	} else {
		fmt.Println(year, "artık yıl değildir.") // Aksi takdirde, artık yıl değildir.
	}

	// ----------------------
	// If Yapısında Initialization
	// ----------------------

	// Eğer if bloğunda bir değişken tanımlanırsa, sadece o blokta geçerli olur.
	// Bu sayede kodu daha güvenli ve düzenli tutabiliriz.
	if number := 10; number > 5 {
		fmt.Println("Number 5'ten büyüktür") // 10 > 5 olduğu için burası yazdırılır.
	} else {
		fmt.Println("Number Küçüktür") // Bu yazdırılmaz.
	}

	// Ekstra Not:
	// Eğer if bloğunda bir değişken tanımlanmazsa, dışarıda bir değişken tanımlanması gerekebilir.
	// Ancak dışarıda tanımlamak daha geniş bir kapsam yaratır, bu da bazen gereksiz veya hatalı kullanımlara yol açabilir.
	// Eğer değişkeni sadece if bloğunda kullanacaksa, initialization yaparak kodu daha temiz ve güvenli hale getirebiliriz.
}
