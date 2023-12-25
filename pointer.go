package main
import "fmt"
func editAge(age *int) {
	*age = *age - 2
}
func pointerFunctions() {
	age:= 164
	agePointer := &age
	// var poi *int = &age
	editAge(&age)
	fmt.Println(age, agePointer)
}