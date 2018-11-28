package main

import "testing"
import db "github.com/iecheniq/go_bootcamp/db"

func TestCrate(t *testing.T) {

	type User struct {
		name    string
		age     int
		salaray float32
	}
	items := []db.Item{
		{"1", User{
			name:    "Ismael Echenique",
			age:     26,
			salaray: 6000.50,
		}},

		{"2", User{
			name:    "Carlos García",
			age:     35,
			salaray: 7500.50,
		}},
	}

	database, _ := db.New("Globant")

	for _, item := range items {
		database.Create("Users", item)
		if _, ok := database.Tables["Users"]; !ok {
			t.Errorf("The table was not created successfully")
		}
		if _, ok := database.Tables["Users"].Items[item.Id]; !ok {
			t.Errorf("The item was not created successfully")
		}
	}
}
