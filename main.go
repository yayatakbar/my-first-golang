package main

import "fmt"

func main() {

	//Number
	fmt.Println("satu =", 1)
	fmt.Println("dua =", 2)
	fmt.Println("tiga koma lima =", 3.5)

	//Boolean
	fmt.Println("True =", true)
	fmt.Println("False =", false)

	//String
	fmt.Println("Hello World")
	fmt.Println(len("Hello World"))
	fmt.Println(("Hello World"[1]))

	//variabel
	var name string
	name = "Akbar"
	fmt.Println(name)

	name = "Yayat"
	fmt.Println(name)

	var iniString = "ini string"
	fmt.Println(iniString)

	inisialisasi := "hasil inisialisasi"

	fmt.Println(inisialisasi)

	var (
		firstName = "Muhammad Sayuti"
		lastName  = "Akbar"
		angkaFav  = 7
	)

	fmt.Println(firstName + " " + lastName)
	fmt.Println(angkaFav)

	//constant

	const consFirstName string = "Cons Akbar"
	const consLastName = "Cons Sayuti"
	fmt.Println(consFirstName)
	fmt.Println(consLastName)

}
