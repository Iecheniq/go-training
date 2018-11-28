package main

import "testing"
import db "github.com/iecheniq/go_bootcamp/db"

func TestDelete(t *testing.T) {

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
	database.Delete("Users", "1")
	if _, ok := database.Tables["Users"].Items["1"]; ok {
		t.Errorf("The item was not deleted successfully")
	}
}
func TestDeleteTable(t *testing.T) {

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
	database.DeleteTable("Users")
	if _, ok := database.Tables["Users"]; ok {
		t.Errorf("The table was not deleted successfully")
	}
}
