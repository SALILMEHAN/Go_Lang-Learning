package main

import (
	"abc/any_name"
	"fmt"
)

// https://www.youtube.com/playlist?list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm

func main() {
	// Video 5:------------------------------------------------------------
	// go mod init abc
	fmt.Println("Hi!")
	// go run main.go ---->> execute this on terminal

	// Video 6:------------------------------------------------------------
	any_name.PrintName("Salil")

	// Video 7:------------------------------------------------------------
	var name string = "yy" //declating like this shows the variable is 100% string
	var version = "ee"     //declaring like this, go lang compiler automatically detects the type of variable declared
	fmt.Println(name)
	fmt.Println(version)

	var money = 1234
	fmt.Println("Price:", money)

	var decided = true
	fmt.Println(decided)

	const pi = 3.14
	// pi = 3		//it will show error
	fmt.Println(pi)

	//This is important, for declaring the variable most developers use like below
	myName := "Salil"
	fmt.Println(myName)

	var PublicVariable = 123 //when we initialize the variable name with 1st letter capital then that variable is exportable for other files. Ex:- Println() function
	var privateVariable = 456
	fmt.Println(PublicVariable)
	fmt.Println(privateVariable)
}
