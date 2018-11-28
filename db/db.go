package db

import (
	"fmt"
)

type Database interface {
	Create(tableName string, items ...Item) error
	Retrieve(tableName string, id string) (Item, error)
	Update(tableName string, id string, item Item) error
	Delete(tableName string, id string) error
	DeleteTable(tableName string)
}

type InMemoryDB struct {
	Name   string
	Tables map[string]Table
}

type Table struct {
	Name  string
	Items map[string]Item
}

type Item struct {
	Id      string
	content interface{}
}

func (db *InMemoryDB) Create(tableName string, items ...Item) error {

	_, ok := db.Tables[tableName]
	if !ok {
		if tableName == "" {
			return fmt.Errorf("Table must have a name ")
		}
		table := Table{
			Name: tableName,
		}
		for _, item := range items {
			if item.Id == "" {
				return fmt.Errorf("Item must have a name")
			}
			table.Items[item.Id] = item
		}
		db.Tables[tableName] = table
		fmt.Printf("Table %v created", tableName)
	} else {
		for _, item := range items {
			if item.Id == "" {
				return fmt.Errorf("Item must have an ID")
			}
			db.Tables[tableName].Items[item.Id] = item
		}
		fmt.Printf("Items created in table %v", tableName)
	}
	return nil
}

func (db *InMemoryDB) Retrieve(tableName string, id string) (Item, error) {

	_, ok := db.Tables[tableName]
	if !ok {
		return Item{}, fmt.Errorf("Table %v does not exist", tableName)
	}

	element, ok := db.Tables[tableName].Items[id]

	if !ok {
		return element, fmt.Errorf("Element with ID %v does not exist", id)
	}
	return element, nil
}

func (db *InMemoryDB) Update(tableName string, id string, item Item) error {

	_, ok := db.Tables[tableName]
	if !ok {
		return fmt.Errorf("Table %v does not exist", tableName)
	}
	if _, ok := db.Tables[tableName].Items[id]; !ok {
		return fmt.Errorf("Element with ID %v does not exist", id)
	}
	if item.Id == "" {
		return fmt.Errorf("Item must have an ID")
	}
	db.Tables[tableName].Items[id] = item
	fmt.Printf("Table %v updated", tableName)
	return nil
}

func (db *InMemoryDB) Delete(tableName string, id string) error {
	_, ok := db.Tables[tableName]
	if !ok {
		return fmt.Errorf("Table %v does not exist", tableName)
	}
	if _, ok := db.Tables[tableName].Items[id]; !ok {
		return fmt.Errorf("Element with ID %v does not exist", id)
	}
	delete(db.Tables[tableName].Items, id)
	fmt.Printf("Item %v deleted in table %v", id, tableName)
	return nil
}

func (db *InMemoryDB) DeleteTable(tableName string) error {
	_, ok := db.Tables[tableName]
	if !ok {
		return fmt.Errorf("Table %v does not exist", tableName)
	}
	delete(db.Tables, tableName)
	fmt.Printf("Table %v deleted", tableName)
	return nil
}

func New(name string) (InMemoryDB, error) {

	db := InMemoryDB{
		Name:   name,
		Tables: make(map[string]Table),
	}
	if name == "" {
		err := fmt.Errorf("Database must have a name")
		return db, err
	}
	return db, nil
}
