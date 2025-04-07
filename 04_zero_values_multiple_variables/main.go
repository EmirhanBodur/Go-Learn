package main

import "fmt"

func main() {
	// var name string = "Emirhan"   
	// var age int = 22
	// var isMarried bool = true
	// var weight float32 = 72.5

	//Yukarıdaki gibi her değişkene var tanımlaması yapmak yerine tüm değişkenleri tek bir "var" bloğuna toplayarak kodumuzu daha temiz, düzenli ve anlaşılabilir yazabiliriz.
	//Ayrıca her değişkenin tipi ve değeri açıkça belli olur, özellikle birden fazla veriyle çalışırken kod daha temiz olur
	/* var (
		name string = "Emirhan"
		age int = 22
		isMarried bool = true
		weight float32 = 72.5
	) */

	// Bu yöntemi kullanırız çünkü tüm değişkenleri tek satırda tanımlayıp değer atayabiliriz
	// Yukarıdakine nazaran hızlı ve pratik yöntemdir. küçük ve basit veri gruplarında kullanılır.
	// Go tipi değerlerden otomatik olarak algılar, bu sayede ayrıca tip belirtmeye gerek kalmaz.
	var name, age, isMarried, weight, height = "Emirhan", 22, true, 22.5, 172
    /* ZERO VALUES
	var a int // 0
	var b string // ""
	var isOnline bool // otomatikmen "false" değeri atar
	*/

	fmt.Println(name, age, isMarried,weight, height)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

}