package main

import "testing"
import db "github.com/iecheniq/go_bootcamp/db"

func TestRetrieve(t *testing.T) {

	type User struct {
		name    string
		age     int
		salaray float32
	}
	item := db.Item{
		Id: "1",
		Content: User{
			name:    "Ismael Echenique",
			age:     26,
			salaray: 6000.50,
		},
	}
	database, _ := db.New("Globant")
	database.Create("Users", item)
	if element, _ := database.Retrieve("Users", "1"); element.Id != "1" {
		t.Errorf("Could not retrieve item")
	}
}
