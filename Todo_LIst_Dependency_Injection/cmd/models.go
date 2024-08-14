/* models.go */

package main

import (
	"database/sql"
	"log"
)

type Todo struct {
	Id        int
	Item      string
	Completed int
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) GetAllTodos() ([]Todo, error) {
	rows, err := m.DB.Query("SELECT * FROM todos")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Item, &todo.Completed)
		if err != nil {
			log.Println(err)
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	return todos, err
}

func (m *TodoModel) Insert(item string, completed int) error {
	stmt := "INSERT INTO todos (item, completed) VALUES (?, ?)"
	_, err := m.DB.Exec(stmt, item, completed)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (m *TodoModel) Update(id int) error {
	stmt := "UPDATE todos SET completed = 1 WHERE id = ?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (m *TodoModel) Delete(id int) error {
	stmt := "DELETE FROM todos WHERE id = ?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

