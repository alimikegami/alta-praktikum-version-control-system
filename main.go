package main

import "fmt"

// function to greet the user
func greet() {
	fmt.Println("Hi, my name is Alim Ikegami! Nice to meet you!")
}

// function to ask question to the user
func askQuestion() {
	fmt.Println("How are you today?")
}

// the main function
func main() {
	fmt.Println("Hello, World!")
	greet()
	askQuestion()
}
