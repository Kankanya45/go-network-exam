package main

import "fmt"

func main() {
	// TODO: Implement a simple Go program
	fmt.Println(HelloFunction())
}

func HelloFunction() string {
	return "Hello"
}
func Add(a int, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func UserLogin(username string, password string) bool {
	if username == "admin" && password == "password" {
		return true
	} else {
		return false
	}
}

// Use function UserLogin
func LogIn(username string, password string) {
	if UserLogin(username, password) {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}
}
