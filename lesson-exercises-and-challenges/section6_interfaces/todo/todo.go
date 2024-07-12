package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) Todo {
	return Todo{
		Text: text,
	}
}

func (toDo Todo) DisplayState() {
	fmt.Println(toDo.Text)
}

func (toDo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(toDo)
	if err != nil {
		return err
	}

	os.WriteFile(filename, json, 0644)
	return nil
}
