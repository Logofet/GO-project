package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) getAllUsers() ([]User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("db.Exec: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *UserModel) addUser(user User) error {
	_, err := u.DB.Exec("INSERT INTO users (id, name, email, age) VALUES (0, ?, ?, ?)", user.Name, user.Email, user.Age)
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}

	return nil
}

func (u *UserModel) updateUser(user User, id string) error {
	_, err := u.DB.Exec("UPDATE users SET name=?, email=?, age=? WHERE id=?", user.Name, user.Email, user.Age, id)
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}

	return nil
}

func (u *UserModel) deleteUser(id string) error {
	_, err := u.DB.Exec(fmt.Sprintf("DELETE FROM users WHERE ID = %s", id))
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}

	return nil
}
