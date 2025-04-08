// Constant nedir ?
// Bir programın çalışma süresi boyunca değişmeyen veri değerine constant diyoruz. Kod içine const ile belirtiyoruz.

package main

import "fmt"

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	const pi = 3.14
	var r = 5

	// Alan hesaplama : pi * r * r
	var area = pi * float64(r*r)
	fmt.Println("Dairenin alanı:", area)

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)

	// Neden sabit(const) kullanmalıyız ?
	// Çünkü sabitler:
	// - Değeri değişmemeli olan bilgileri korur (örneğin pi, sabit vergi oranı vs.)
	// - Hatalı değişimleri engeller
	// - Kodun okunabilirliğini ve güvenilirliğini artırır

}
