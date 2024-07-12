package main

/*
	Experimenting with structs and custom data types.
*/

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstname string
	lastname  string
	birthDate string
	createdAt time.Time
}

// Embedding structs - something like Inheritance in Java.
type admin struct {
	email    string
	password string
	user
}

func (u *user) changeUserInfo() {
	fmt.Print("What would you like to change: firstname, lastname or birthDate? For exit type 'exit'\nType one of the above choices: ")

	userChoice := getSpecificUserInput()

	switch userChoice {
	case "firstname":
		fmt.Print("Provide a new firstname: ")
		u.firstname = getSpecificUserInput()
	case "lastname":
		fmt.Print("Provide a new lastname: ")
		u.lastname = getSpecificUserInput()
	case "birthday":
		fmt.Print("Provide a new birth date: ")
		u.birthDate = getSpecificUserInput()
	case "exit":
		fmt.Print("exiting...")
		return
	default:
		fmt.Println("wrong input")
		return
	}

}

// newUser is a Creation/Constructor function. It also validates based on specific criteria
func newUser(firstname, lastname, birthday string) (user, error) {
	if firstname == "" || lastname == "" || birthday == "" {
		return user{}, errors.New("firstname, lastname and birthday are required")
	}

	return user{
		firstname: firstname,
		lastname:  lastname,
		birthDate: birthday,
		createdAt: time.Now(),
	}, nil
}

func newAdmin(email string, password string) admin {
	return admin{
		email:    email,
		password: password,
		user: user{
			firstname: "Panos",
			lastname:  "Vasilopoulos",
			birthDate: "------",
			createdAt: time.Now(),
		},
	}
}

func main() {
	// userFirstname, userLastname, userBirthday := getUserInitialInfo()

	// appUser1 := user{
	// 	firstname: userFirstname,
	// 	lastname:  userLastname,
	// 	birthDate: userBirthday,
	// 	createdAt: time.Now(),
	// }

	user2Firstname, user2Lastname, user2Birthday := getUserInitialInfo()

	appUser2, err := newUser(user2Firstname, user2Lastname, user2Birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	// printUserInfo(appUser1)

	printUserInfo(appUser2)

	for counter := 0; counter < 1; counter++ {
		appUser2.changeUserInfo()
	}

	printUserInfo(appUser2)

	admin := newAdmin("panos@gmail.com", "testing1234")
	printAdminInfo(admin)
}

func printUserInfo(u user) {
	fmt.Println(u.firstname, u.lastname, u.birthDate, u.createdAt.Format("2006-01-02 15:04:05"))
}

func printAdminInfo(adm admin) {
	fmt.Println(adm.email, adm.password, adm.firstname, adm.lastname, adm.birthDate, adm.createdAt.Format("2006-01-02 15:04:05"))
}

func getUserInitialInfo() (string, string, string) {
	var firstname string
	var lastname string
	var birthday string

	fmt.Print("Please provide your firstname: ")
	errorInUserInfo := errors.New("error in user input")

	_, err := fmt.Scan(&firstname)
	if err != nil {
		fmt.Println(errorInUserInfo)
	}

	fmt.Print("Please provide your lastname: ")
	_, err = fmt.Scan(&lastname)
	if err != nil {
		fmt.Println(errorInUserInfo)
	}

	fmt.Print("Please provide your birthday: ")
	_, err = fmt.Scan(&birthday)
	if err != nil {
		fmt.Println(errorInUserInfo)
	}

	return firstname, lastname, birthday
}

func getSpecificUserInput() string {
	var userInfo string
	_, err := fmt.Scanf("%v\n", &userInfo)
	if err != nil {
		fmt.Println(err)
	}

	return userInfo
}
