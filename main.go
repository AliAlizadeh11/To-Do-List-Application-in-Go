package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("welcome to ToDo app ")

	command := flag.String("command", "no command", "command to run")
	flag.Parse()

	for {
		runcommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Printf("userstorage: %+v\n@", userStorage)
}

func runcommand(command string) {
	switch command {

	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)

	}

}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("task:", name, duedate, category)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category:", title, color)

}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", email, password)

}
