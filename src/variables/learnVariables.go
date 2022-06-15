package main

import (
	"fmt"
	"strconv"
)

var i int = 42

var (
//aa string = "dude"
//vv int    = 2
)

func main() {

	//i := 42
	fmt.Printf("%v, %T", i, i)
	println("\n_____________")
	//fmt.Println(aa, vv)
	var j float32 = float32(i)
	fmt.Printf("%v, %T", j, j)
	println("\n_____________")

	var toString string = strconv.Itoa(i)
	fmt.Printf("%v, %T", toString, toString)

	println("\n_____________")
	boolValue := true
	fmt.Printf("%v, %T", boolValue, boolValue)

	println("\n_____________")
	myString := "the string"
	fmt.Printf("%v, %T", myString[2], myString[2])

	println("\n_____________")
	bytes := []byte(myString)
	fmt.Printf("%v, %T", bytes, bytes)

	println("\n_____________")
	grades := [...]int{10, 2, 3, 4} // ... make just enough allocation to fit data
	gradesDuplicate := &grades      //this points to same address as grades, otherwise copies values
	fmt.Printf("%v", grades)
	fmt.Printf("%v", gradesDuplicate)

	println("\n_____________")
	var students = []string{"dasd", "bb", "cc"} //called slice, almost the same as List (array backed, but unlimited size), works as pointers
	fmt.Printf("%v", students)
	fmt.Printf("%v, %v", students[0], len(students))

	println("\n_____________")
	students = append(students, "ada", "dasda", "john")
	fmt.Printf("%v, %v", students, len(students))

	println("\n_____________")
	newMap := make(map[string]int)
	newMap = map[string]int{"key1": 1111}
	fmt.Println(newMap)

}
