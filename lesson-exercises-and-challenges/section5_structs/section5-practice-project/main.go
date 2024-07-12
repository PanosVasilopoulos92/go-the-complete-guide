package main

import (
	"bufio"
	"fmt"
	"os"
	"section5-practice-project/notebook"
)

func main() {
	notebookName, data := getUserInput()

	notebook1 := notebook.New(notebookName, data)
	fmt.Println(notebook1)

	// notebook1.AddNotes("jvw243692vks")
	fmt.Println(notebook1.GetData())

	notebook1.DisplayState()

	err := notebook1.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File successfully created.")

}

func getUserInput() (string, string) {
	var name string
	var data string

	fmt.Print("Hi! what name you wanna give to your notebook? ")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Write what you have in your mind:")
	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	data += line
	// 	if line == "" {
	// 		break
	// 	}
	// }

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	// input = strings.TrimSuffix(input, "\n")
	// input = strings.TrimSuffix(input, "\r")

	data += input
	return name, data
}
