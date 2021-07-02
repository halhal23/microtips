package repository

import (
	"context"
	"database/sql"
	"microtips/user/pb"

	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	InsertUser(ctx context.Context, input *pb.UserInput) (int64, error)
	SelectUserById(ctx context.Context, id int64) (*pb.User, error)
	UpdateUser(ctx context.Context, id int64, input *pb.UserInput) error
	DeleteUser(ctx context.Context, id int64) error
	SelectAllUsers() (*sql.Rows, error)
}

type sqliteRepo struct {
	db *sql.DB
}

func NewsqliteRepo() (Repository, error) {
	db, err := sql.Open("sqlite3", "./user/user.sql")
	if err != nil {
		return nil, err
	}
	// usersテーブルを作成
	cmd := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		password STRING)`

	_, err = db.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return &sqliteRepo{db}, nil
}

func (r *sqliteRepo) InsertUser(ctx context.Context, input *pb.UserInput) (int64, error) {
	cmd := "INSERT INTO users(name, password) VALUES (?, ?)"
	result, err := r.db.Exec(cmd, input.Name, input.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *sqliteRepo) SelectUserById(ctx context.Context, id int64) (*pb.User, error) {
	cmd := "SELECT * FROM users WHERE id = ?"
	row := r.db.QueryRow(cmd, id)
	var user pb.User
	err := row.Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:       int64(user.Id),
		Name:     user.Name,
		Password: user.Password,
	}, nil
}

func (r *sqliteRepo) UpdateUser(ctx context.Context, id int64, input *pb.UserInput) error {
	cmd := "UPDATE users SET name = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(cmd, input.Name, input.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteRepo) DeleteUser(ctx context.Context, id int64) error {
	cmd := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteRepo) SelectAllUsers() (*sql.Rows, error) {
	cmd := "SELECT * FROM users"
	rows, err := r.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
