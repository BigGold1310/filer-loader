package models

import (
	"database/sql"
	"errors"
)

// Task - Struct to pass tasks through the goroutines
type Task struct {
	ID       int            `db:"id"`
	Link     string         `db:"link"`
	FileName string         `db:"name"`
	FileHash string         `db:"hash"`
	Group    int            `db:"fk_group"`
	Status   int            `db:"fk_status"`
	Message  sql.NullString `db:"message"`
	Size     int            `db:"size"`
}

// GetAllTasks - Returns all tasks from the database
func (db *DB) GetAllTasks() ([]Task, error) {
	rows, err := db.Queryx("SELECT * FROM tbl_tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		err = rows.StructScan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTasksByStatus - Returns a filtered task list by status
func (db *DB) GetTasksByStatus(status int) ([]Task, error) {
	rows, err := db.Queryx("SELECT * FROM tbl_tasks WHERE status == $1", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		err = rows.StructScan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTasksByStatusLimit - Returns a filtered task list by status
func (db *DB) GetTasksByStatusLimit(status int, limit int) ([]Task, error) {
	if limit > 0 {
		rows, err := db.Queryx("SELECT * FROM tbl_tasks WHERE status == $1 LIMIT $2", status, limit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		tasks := make([]Task, 0)
		for rows.Next() {
			var task Task
			err = rows.StructScan(&task)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		return tasks, nil
	} else if limit <= 0 {
		return nil, errors.New("limit must greater than 0")
	}
	// dummy return to make linter happy
	return nil, nil
}

// AddTask - Returns a filtered task list by status
func (db *DB) AddTask(status int) ([]Task, error) {
	rows, err := db.Queryx("INSERT * FROM tbl_tasks WHERE tbl_tasks.status == $1", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		err = rows.StructScan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateTask - Returns a filtered task list by status
func (db *DB) UpdateTask(task Task) ([]Task, error) {
	rows, err := db.Queryx("SELECT * FROM tbl_tasks WHERE tbl_tasks.status == $1", task)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		
		err = rows.StructScan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask - Deletes a task
func (db *DB) DeleteTask(id int) error {
	err := deleteTask(db, id)
	if err != nil {
		return err
	}
	return nil
}

// deleteTask - Deletes a task (effective delete function)
func deleteTask(db *DB, id int) error {
	_, err := db.Exec("DELETE FROM tbl_tasks WHERE tbl_tasks.id == $1", id)
	if err != nil {
		return err
	}
	// ToDo:
	// - Log the result of the deletion
	return nil
}
