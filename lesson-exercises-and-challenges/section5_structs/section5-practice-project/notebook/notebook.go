package notebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type NoteBook struct {
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func New(name, data string) NoteBook {
	if name == "" || data == "" {
		fmt.Println("name and data cannot be empty")
	}

	return NoteBook{
		Name:      name,
		Data:      data,
		CreatedAt: time.Now(),
	}
}

func (n NoteBook) DisplayState() {
	fmt.Printf("Notebook '%v' created at %v and has the following content so far:\n%v",
		n.Name, n.CreatedAt.Format("2006-01-02"), n.Data)
}

func (n NoteBook) Save() error {
	filename := strings.ReplaceAll(n.Name, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
		return errors.New("error while marshaling")
	}

	os.WriteFile(filename, json, 0644)
	return nil
}

func (n *NoteBook) AddNotes(note string) {
	n.Data += note
}

func (d *NoteBook) GetData() string {
	return d.Data
}
