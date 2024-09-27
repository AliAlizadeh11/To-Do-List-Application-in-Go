package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserId     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var userStorage []User
var authenticatedUser *User

var taskStorage []Task
var categoryStorage []Category

const userStoragePath = "user.txt"

func main() {

	loadUserStorageFromFile()
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

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("category id is not valid")

		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Println("category id is not a valid integer, %v\n", err)

		return
	}

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserId:     authenticatedUser.ID,
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

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)

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

	writeUserToFile(user)

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

func loadUserStorageFromFile() {
	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Println("can't open file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read file", err)
	}

	var dataStr = string(data)

	dataStr = strings.Trim(dataStr, "\n")

	userSlice := strings.Split(dataStr, "\n")
	for _, u := range userSlice {
		if u == "" {
			continue
		}

		var user = User{}

		userFields := strings.Split(u, ",")

		for _, field := range userFields {
			values := strings.Split(field, ": ")
			if len(values) != 2 {
				fmt.Println("invalid user field, skipping...", len(values))

				continue
			}

			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					fmt.Println("can't convert id to int", err)

					return
				}
				user.ID = id

			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue

			}

			fmt.Printf("user: %+v\n", user)
		}
	}
	fmt.Println(data)
}

func writeUserToFile(user User) {

	var file *os.File

	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can't create or open file", err)

		return
	}

	defer file.Close()
	data := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name, user.Email, user.Password)

	var b = []byte(data)

	numberOfWritten, wErr := file.Write(b)

	if wErr != nil {
		fmt.Println("can't write to the file %v\n", wErr)

		return
	}

	fmt.Printf("wrote %d bytes\n", numberOfWritten)
}
