package main
import (
	"fmt"
	"time"
)

type User struct { // will be available outside also
	firstName string
	lastName string
	createdAt time.Time
}

// we can create a constructor type
func newUser(firstName, lastName string, createdAt time.Time)  *User{
	
	return &User {
		firstName: firstName,
		lastName: lastName,
		createdAt: createdAt,
	}
}
func structres() {
	var appUser *User
	appUser = newUser("Lokesh", "Naik", time.Now())


	var user1 User
	user1 = User{
		firstName: "Lavadiya Lokesh",
		lastName: "Naik",
		createdAt: time.Now(),
	}
	outputUserDetails(&user1)
	// always pointer is passed here even if we don't specify

	appUser.structureFunction()
	// a reference (copy) is created when this is called
	// copy will be called

	outputUserDetails(appUser)
}

func outputUserDetails(u *User) {
	fmt.Println((*u).firstName, u.lastName, u.createdAt)
}


func (u *User) structureFunction() {
	// (u user) is called receiver parameter
	u.firstName = "Lavadiya"
	// structureFunction this is a function of the structure
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}