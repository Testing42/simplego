package main

import "fmt"

func main() {
	age := 32 //regular variable no pointer

	var agePointer *int
	// this is not neeeded but makes it clear a pointer is being created with the * in front of the int
	//agePointer := &age //how to create a pointer pointing to the age memory address.
	agePointer = &age // since the variable was initialized you don't need the

	fmt.Println("Age:", *agePointer) //this * gets the value behind the pointer

	editAgeToAdultYears(agePointer)

	fmt.Println(age)
}

/*func getAdultYears(age int) int {
   function before pointers
   return age - 18
}
*/

func editAgeToAdultYears(age *int) {
	/*you can't perform calculations on a
	pointers. You will have to do dereferencing
	 to the pointer
	*/
	//return *age - 18
	/*To save on calling the pointer you can dereference the
	pointer so the actual value is changed and less memory is used.
	This saves memory because you put the result back in the same memory
	and not use any additional memory space.

	This now overides the 32 value at the top of the script
	with the value below

	This means we need to remove the int as getAdultYears will no longer return anything
	*/
	*age = *age - 18
}
