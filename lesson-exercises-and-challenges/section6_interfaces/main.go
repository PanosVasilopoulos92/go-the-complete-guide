package main

import (
	"bufio"
	"fmt"
	"lesson-exercises-and-challenges/section6_interfaces/notebook"
	"lesson-exercises-and-challenges/section6_interfaces/todo"
	"os"
)

type saver interface {
	Save() error
}

func main() {
	userText := getUserData("Write anything you wish: ")
	text1 := todo.New(userText)
	fmt.Println(text1.Text)
	// text1.Save()
	saveData(text1)
	fmt.Println("-------")

	title := getUserData("Enter notebook's title: ")
	content := getUserData("Write down your note: ")
	notebook1 := notebook.New(title, content)
	// notebook1.Save()
	saveData(notebook1)
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return userInput
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Data was not saved.")
		return err
	} else {
		fmt.Println("Data successfully saved.")
		return nil
	}
}
