package main

import "testing"
import db "github.com/iecheniq/go_bootcamp/db"

func TestUpdate(t *testing.T) {

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

	item = db.Item{
		Id: "2",
		Content: User{
			name:    "Ismael Echenique",
			age:     26,
			salaray: 6000.50,
		},
	}
	database.Update("Users", "1", item)
	if database.Tables["Users"].Items["1"].Id != "2" {
		t.Errorf("The item was not updated successfully")
	}
}
