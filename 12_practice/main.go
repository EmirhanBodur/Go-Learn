/* Soru 1: Typeless Constant & Type Conversion
package main

import "fmt"
Bir int16 türünde bir değişken tanımla. Daha sonra bir typeless constant (örneğin const pi = 3.14) tanımla. Bu iki değeri çarp ve sonucu ekrana yazdır.
Ama dikkat: Tür farklılığından dolayı type conversion gerekebilir. Bu detayı unutma.
func main() {
	const pi = 3.14
	 1. yöntem
	var number int16 = 15
	fmt.Println(pi * float32(number))
	const number =15
	fmt.Println(pi*number)
}
*/
/* Soru 2: Shadowing Global scope'ta bir değişken tanımla (örneğin var name = "Global").
Sonra main fonksiyonu içinde aynı isimli ama farklı bir değer içeren başka bir name değişkeni tanımla.
Her iki name değişkenini doğru şekilde yazdırarak shadowing kavramını göster.


package main

import "fmt"

var name = "Global"

func main() {

	name := "Local"
	fmt.Println(name)
	globalVar()

}

func globalVar() {
	fmt.Println(name)
}
*/
/*
SORU 3: İki farklı int değişken tanımla ve bu değişkenlerle toplama, çıkarma, çarpma, bölme, mod alma işlemleri yap.
Sonuçları ekrana açıklayıcı metinle birlikte yazdır. (fmt.Printf ile de olabilir.)
package main

import "fmt"

func main() {
	var number1 int = 32
	var number2 int64 = 64

	toplam := int64(number1) + number2
	cikar := int64(number1) - number2
	carp := int64(number1) * number2
	bol := int64(number1) / number2
	mod := int64(number1) % number2

	fmt.Printf("Toplam: %d\n", toplam)
	fmt.Printf("Çıkarma: %d\n", cikar)
	fmt.Printf("Çarpma: %d\n", carp)
	fmt.Printf("Bölme: %d\n", bol)
	fmt.Printf("Mod: %d\n", mod)
}
*/
package main

import "fmt"

func main() {
	fmt.Println(" /$$   /$$           /$$ /$$                  /$$$$$$  /$$   /$$     /$$                 /$$      ")
	fmt.Println("| $$  | $$          | $$| $$                 /$$__  $$|__/  | $$    | $$                | $$      ")
	fmt.Println("| $$  | $$  /$$$$$$ | $$| $$  /$$$$$$       | $$  \\__/ /$$ /$$$$$$  | $$$$$$$  /$$   /$$| $$$$$$$ ")
	fmt.Println("| $$$$$$$$ /$$__  $$| $$| $$ /$$__  $$      | $$ /$$$$| $$|_  $$_/  | $$__  $$| $$  | $$| $$__  $$")
	fmt.Println("| $$__  $$| $$$$$$$$| $$| $$| $$  \\ $$      | $$|_  $$| $$  | $$    | $$  \\ $$| $$  | $$| $$  \\ $$")
	fmt.Println("| $$  | $$| $$_____/| $$| $$| $$  | $$      | $$  \\ $$| $$  | $$ /$$| $$  | $$| $$  | $$| $$  | $$")
	fmt.Println("| $$  | $$|  $$$$$$$| $$| $$|  $$$$$$/      |  $$$$$$/| $$  |  $$$$/| $$  | $$|  $$$$$$/| $$$$$$$/")
	fmt.Println("|__/  |__/ \\_______/|__/|__/ \\______/        \\______/ |__/   \\___/  |__/  |__/ \\______/ |_______/      ")
}
