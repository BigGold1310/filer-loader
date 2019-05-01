package models

// Datastore - Datastore holds all functions which should be exported
type Datastore interface {
	AllBooks() ([]*Book, error)
}
