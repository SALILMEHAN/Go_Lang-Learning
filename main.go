package main

import (
	"fmt"
	"time"
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

func modifyingDataThroughPassingReference(value *int) {
	*value = *value * 5
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
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(ans)
	// }

	// Video 12 (arrays):------------------------------------------------------------
	// var number [5]int
	// number[0] = 1
	// number[1] = 2
	// number[2] = 3
	// number[3] = 4
	// number[4] = 5
	// fmt.Println(number[0])

	// var num = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(num[3])

	// fmt.Println("Length of num is:", len(num))

	// var prices [10]int //after initialization of array every value assigned will be 0
	// fmt.Println(prices)

	// Video 13 (slices):------------------------------------------------------------
	// var numbers = make([]int, 3, 5) // Type, Length, Capacity

	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// numbers = append(numbers, 4) //when the length==capacity, then we append the element the capacity doubles
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// Video 14 (If-Else):------------------------------------------------------------
	// age := 25
	// if age > 18 {
	// 	fmt.Println("You are an adult")
	// } else if age == 18 {
	// 	fmt.Println("You are a 18")
	// } else {
	// 	fmt.Println("You are a minor")
	// }

	// Video 15 (Switch-Case):------------------------------------------------------------
	// day := "Monday"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("Weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// temperature := 25
	// switch {
	// case temperature > 30:
	// 	fmt.Println("It's hot")
	// case temperature < 10:
	// 	fmt.Println("It's cold")
	// default:
	// 	fmt.Println("It's normal")
	// }

	// Video 16 (For Loop):------------------------------------------------------------
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// counter := 0
	// for {
	// 	fmt.Println(counter)
	// 	counter++
	// 	if counter == 5 {
	// 		break
	// 	}
	// }

	// numbers := []int{11, 12, 13, 14, 15}
	// for ind, val := range numbers {
	// 	fmt.Println(ind, val)
	// }

	// Video 17 (Maps):------------------------------------------------------------
	// marks := make(map[string]int)

	// marks["Alice"] = 80
	// marks["Bob"] = 90
	// marks["Charlie"] = 70

	// fmt.Println(marks["Alice"])

	// delete(marks, "Bob")
	// fmt.Println(marks["Bob"])

	// //check for key present or not
	// grade, exist := marks["Alice"]
	// if exist {
	// 	fmt.Println("Grade of Alice is ", grade)
	// } else {
	// 	fmt.Println("No grade of Alice")
	// }

	// grade2, exist2 := marks["David"]
	// if exist2 {
	// 	fmt.Println("Grade of David is ", grade2)
	// } else {
	// 	fmt.Println("No grade of David")
	// }

	// for ind, val := range marks {
	// 	fmt.Println(ind, val)
	// }

	// Video 18 (Struct):------------------------------------------------------------

	// type Person struct {
	// 	FirstName string
	// 	LastName  string
	// 	Age       int
	// }

	// //declaration type 1
	// var salil Person

	// salil.FirstName = "SALIL"
	// salil.LastName = "MEHAN"
	// salil.Age = 22

	// fmt.Println(salil.FirstName, salil.LastName, salil.Age)
	// fmt.Println(salil)

	// //declaration type 2
	// person1 := Person{
	// 	FirstName: "abc",
	// 	LastName:  "def",
	// 	Age:       23,
	// }

	// fmt.Println(person1.FirstName, person1.LastName, person1.Age)

	// //declaration type 3
	// var person2 = new(Person)

	// person2.FirstName = "grcvcd"
	// person2.LastName = "grvfdf"
	// person2.Age = 54

	// fmt.Println(person2.FirstName, person2.LastName, person2.Age)

	// //Important, struct in struct

	// type Contact struct {
	// 	Phone string
	// 	Email string
	// }

	// type employee struct {
	// 	employee_detail  Person
	// 	employee_contact Contact
	// }

	// var emp1 employee
	// emp1.employee_detail.FirstName = "SALIL"
	// emp1.employee_detail.LastName = "MEHAN"
	// emp1.employee_detail.Age = 22
	// emp1.employee_contact.Phone = "1234567890"
	// emp1.employee_contact.Email = "salilmehan123@gmail.com"

	// fmt.Println(emp1)

	// Video 19 (Pointers):------------------------------------------------------------
	// num := 2
	// ptr := &num
	// fmt.Println(*ptr)
	// fmt.Println(ptr)

	// value := 5
	// modifyingDataThroughPassingReference(&value)
	// fmt.Println(value)

	// Video 20 (Data Conversion):------------------------------------------------------------
	// var a int = 10
	// var b float64 = float64(a)
	// fmt.Println(b)
	// var c int = int(b)
	// fmt.Println(c)
	// var e string = strconv.Itoa(a)
	// fmt.Println(e)
	// f, _ := strconv.ParseFloat("10.5", 64)
	// fmt.Println(f)
	// var g int = int(f)
	// fmt.Println(g)

	// Video 21 (String Packages):------------------------------------------------------------
	// data := "apple,orange,banana"
	// final := strings.Split(data, ",") //Split Function
	// fmt.Println(final)

	// str := "one two three four two five two"
	// out2 := strings.Count(str, "two") //Count Characters
	// fmt.Println(out2)

	// str2 := "               lets go              "
	// out3 := strings.TrimSpace(str2) //Trim
	// fmt.Println(out3)

	// str3 := "Salil"
	// str4 := "Mehan"
	// out4 := strings.Join([]string{str3, str4, "Hello"}, " ") //concatinate
	// fmt.Println(out4)

	// Video 22 (Time Package):------------------------------------------------------------
	t := time.Now()
	fmt.Println(t)

	formatted := t.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println(formatted)

	datestr := "2023-11-25"
	layout_str := "2006-01-02"
	formatted_time, _ := time.Parse(layout_str, datestr)
	fmt.Println(formatted_time)

	new_t := t.Add(24 * time.Hour) // adding 1 day
	fmt.Println(new_t)
	new_t1 := new_t.Format("02-01-2006 Monday")
	fmt.Println(new_t1)

}
