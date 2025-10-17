package main

import "fmt"

func main() {
	age := 32 //regular variable no pointer

	//var agePointer *int // this is not neeeded but makes it clear a pointer is being created with the * in front of the int
	agePointer := &age //how to create a pointer pointing to the age memory address.
	//agePointer = &age // since the variable was initialized you don't need the

	fmt.Println("Age:", *agePointer) //this * gets the value behind the pointer

	adultYears := getAdultYears(agePointer)

	fmt.Println(adultYears)
}

/*func getAdultYears(age int) int {
   function before pointers
   return age - 18
}
*/

func getAdultYears(age *int) int {
	/*you can't perform calculations on a
	pointers. You will have to do dereferencing
	 to the pointer
	*/
	return *age - 18
}
