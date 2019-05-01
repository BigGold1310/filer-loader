package models

import "database/sql"

// Task - Struct to pass tasks through the goroutines
type Task struct {
	ID       int
	Link     string
	FileName string
	FileHash string
	Group    int
	Status   int
	Message  sql.NullString
	Size     int
}

// AllBooks -
func (db *DB) AllBooks() ([]*Task, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Task, 0)
	for rows.Next() {
		bk := new(Task)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
