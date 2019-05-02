package models

// Datastore - Datastore holds all functions which should be exported
type Datastore interface {
	GetAllTasks() ([]Task, error)
	GetTasksByStatus(status int) ([]Task, error)
	GetTasksByStatusLimit(status int, limit int) ([]Task, error)
}
