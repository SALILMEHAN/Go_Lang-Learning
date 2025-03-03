package main

import (
	"fmt"
)

func myFunction() {
	fmt.Println("simple function")
}

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// https://www.youtube.com/playlist?list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm

func main() {
	// Video 5:------------------------------------------------------------
	// // go mod init abc
	fmt.Println("Hi!")
	// // go run main.go ---->> execute this on terminal

	// Video 6 (Importing Packages):------------------------------------------------------------
	// any_name.PrintName("Salil")

	// Video 7 (Variables):------------------------------------------------------------
	// var name string = "yy" //declating like this shows the variable is 100% string
	// var version = "ee"     //declaring like this, go lang compiler automatically detects the type of variable declared
	// fmt.Println(name)
	// fmt.Println(version)

	// var money = 1234
	// fmt.Println("Price:", money)

	// var decided = true
	// fmt.Println(decided)

	// const pi = 3.14
	// // pi = 3		//it will show error
	// fmt.Println(pi)

	// //This is important, for declaring the variable most developers use like below
	// myName := "Salil"
	// fmt.Println(myName)

	// var PublicVariable = 123 //when we initialize the variable name with 1st letter capital then that variable is exportable for other files. Ex:- Println() function
	// var privateVariable = 456
	// fmt.Println(PublicVariable)
	// fmt.Println(privateVariable)

	// Video 8 (println and printf in Golang):------------------------------------------------------------
	// age := 22
	// name := "ww"
	// height := 5.6794838
	// fmt.Println("Name:", name, "Age:", age, "Height:", height)

	// fmt.Printf("My height is %.3f", height)
	// fmt.Printf("My age is %d\n", age)
	// fmt.Printf("My name is %s\n", name)
	// fmt.Printf("Name %s, Age %d, Height %.3f", name, age, height)

	// Video 9 (taking inputs):------------------------------------------------------------
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// name, _ := reader.ReadString('\n')
	// fmt.Printf("Your name is %s", name)

	// Video 10 (functions):------------------------------------------------------------
	// myFunction()
	// addOutput := add(3, 4)
	// fmt.Println(addOutput)

	// mulOutput := multiply(3, 4)
	// fmt.Println(mulOutput)

	// Video 11 (errors):------------------------------------------------------------
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
