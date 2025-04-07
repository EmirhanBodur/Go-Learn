// 1-) studentName --> Emirhan Bodur, grade --> 22, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

package main

import "fmt"

func main() {
	//1.Yöntem
	/* var studentName string = "Emirhan Bodur"
	var grade int = 22
	var isPassed bool = true; */

	//2.Yöntem
	/* var studentName, grade, isPassed = "Emirhan Bodur", 22, true */

	//3.Yöntem
	/* studentName, grade, isPassed := "Emirhan Bodur", 22, true */

	//farklı bir yöntem (istersek değişken tipinide belirtebiliriz fakat Go kendisi otomatikmen veri tipini belirlediği için ben böyle kullanmayı daha çok tercih ediyorum.)
	var (
		studentName, grade, isPassed = "Emirhan", 22, true
	)

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

}
