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

type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category string
	IsDone   bool
	UserId   int
}

func (u User) print() {
	fmt.Println("User:", u.ID, u.Email, u.Name)

}

var userStorage []User
var authenticatedUser *User

var taskStorage []Task

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

}

func runcommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {

	case "create-task":
		createTask()

	case "create-category":
		createCategory()

	case "register-user":
		registerUser()

	case "login":
		login()

	case "list-task":
		listTask()

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("command is not valid", command)

	}

}

func createTask() {

	scanner := bufio.NewScanner(os.Stdin)

	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		DueDate:  duedate,
		Category: category,
		IsDone:   false,
		UserId:   authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)
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
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is incorrect")

	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.UserId == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
