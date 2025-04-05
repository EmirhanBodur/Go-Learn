// uint8: 0 ile 255 arasında olan işaretsiz 8-bit tam sayılar
// (0 ile 255 arası değer alabilir)
//var x uint8 = 255

// uint16: 0 ile 65535 arasında olan işaretsiz 16-bit tam sayılar
// (0 ile 65535 arası değer alabilir)
//var y uint16 = 65535

// uint32: 0 ile 4294967295 arasında olan işaretsiz 32-bit tam sayılar
// (0 ile 4294967295 arası değer alabilir)
//var z uint32 = 4294967295

// uint64: 0 ile 18446744073709551615 arasında olan işaretsiz 64-bit tam sayılar
// (0 ile 18446744073709551615 arası değer alabilir)
//var w uint64 = 18446744073709551615

// int8: -128 ile 127 arasında olan işaretli 8-bit tam sayılar
// (-128 ile 127 arası değer alabilir)
//var a int8 = -128

// int16: -32768 ile 32767 arasında olan işaretli 16-bit tam sayılar
// (-32768 ile 32767 arası değer alabilir)
//var b int16 = 32767

// int32: -2147483648 ile 2147483647 arasında olan işaretli 32-bit tam sayılar
// (-2147483648 ile 2147483647 arası değer alabilir)
//var c int32 = 2147483647

// int64: -9223372036854775808 ile 9223372036854775807 arasında olan işaretli 64-bit tam sayılar
// (-9223372036854775808 ile 9223372036854775807 arası değer alabilir)
//var d int64 = -9223372036854775808

// float32: IEEE 754 standardına göre 32-bit kayan noktalı sayılar
// (daha düşük hassasiyetli ondalıklı sayılar)
//var e float32 = 3.14

// float64: IEEE 754 standardına göre 64-bit kayan noktalı sayılar
// (daha yüksek hassasiyetli ondalıklı sayılar)
//var f float64 = 3.14159265358979

// complex64: float32 reel ve sanal kısmına sahip karmaşık sayılar
//var g complex64 = complex(1.0, 2.0)

// complex128: float64 reel ve sanal kısmına sahip karmaşık sayılar
//var h complex128 = complex(1.0, 2.0)

// byte: uint8'nin bir takma adı (alias), 8-bit işaretsiz tam sayı
//var i byte = 255

// rune: int32'nin bir takma adı (alias), 32-bit işaretli tam sayı
// Genellikle Unicode karakterlerini temsil etmek için kullanılır
//var j rune = 'A'

package main

import "fmt"

func main() {
	var name = "Emirhan"
	var surname string = "Bodur"
	var c int32 = -2147483647
	age := 25
	var pi = 3.14

	// isMarried değişkeninde tür belirtilmediği için varsayılan olarak false değeri atanır.
	var isMarried  bool 
	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(c)
	fmt.Println(age)
	fmt.Println(pi)
	fmt.Println(isMarried)

	//eğer Printf'de "%T" böyle yazarsak değişken türlerini yazar ama  bitişik ve aynı satırda yazar. Ancak "%T\n" şeklinde yazarsak aynı Println'deki değişken türleri alt alta sıralanır.
	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",surname)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n",c)
	fmt.Printf("%T\n",pi)


	fmt.Printf("%T\n",isMarried)

	

}