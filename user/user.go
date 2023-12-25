package user
import (
	"fmt"
	"time"
)

type User struct { // will be available outside also
	FirstName string
	lastName string
	createdAt time.Time
}

type Admin struct {
	Name string
	User // inheritance
}

func NewAdmin(adminName string) Admin {
	return Admin {
		Name: adminName,
		User: User{
			FirstName: "ADMIN",
			lastName: "ADMIN",
			createdAt: time.Now(),
		},
	}
}
// we can create a constructor type
func newUser(firstName, lastName string, createdAt time.Time)  *User{
	
	return &User {
		FirstName: firstName,
		lastName: lastName,
		createdAt: createdAt,
	}
}
func UserStructres() {
	var appUser *User
	appUser = newUser("Lokesh", "Naik", time.Now())

	var user1 User
	user1 = User{
		FirstName: "Lavadiya Lokesh",
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
	fmt.Println((*u).FirstName, u.lastName, u.createdAt)
}


func (u *User) structureFunction() {
	// (u user) is called receiver parameter
	u.FirstName = "Lavadiya"
	// structureFunction this is a function of the structure
	fmt.Println(u.FirstName, u.lastName, u.createdAt)
}

func (u *User) StructureFunction() {
	// (u user) is called receiver parameter 
	// structureFunction this is a function of the structure
	fmt.Println(u.FirstName, u.lastName, u.createdAt)
}

func (u *User) OutputUserDetails() {
	fmt.Println((*u).FirstName, u.lastName, u.createdAt)
}
