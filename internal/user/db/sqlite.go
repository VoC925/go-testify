package db

import (
	"database/sql"
	"time"

	"github.com/VoC925/go-testify/internal"
	"github.com/VoC925/go-testify/internal/user"

	_ "modernc.org/sqlite"
)

var _ user.Storage = &UserStore{}

// Хранилище
type UserStore struct {
	db *sql.DB
}

// конструктор
func New(db *sql.DB) user.Storage {
	return &UserStore{
		db: db,
	}
}

// Добавление в ДБ
func (u *UserStore) Add(userForAdd *user.User) (int, error) {
	q := `INSERT INTO users (name, login, password_hash, created_at, change_at)
VALUES (:nameVal, :loginVal, :hash, :timeAdd, :timeUpdate)`
	res, err := u.db.Exec(q,
		sql.Named("nameVal", userForAdd.Name),
		sql.Named("loginVal", userForAdd.Login),
		sql.Named("hash", userForAdd.PasswordHash),
		sql.Named("timeAdd", userForAdd.CreatedAt),
		sql.Named("timeUpdate", userForAdd.LastChangedAt))
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Извлечение по логину
func (u *UserStore) GetByLogin(login string) (*user.User, error) {
	q := `SELECT name, login, password_hash, created_at, change_at
FROM users WHERE login=:loginVal`
	row := u.db.QueryRow(q,
		sql.Named("loginVal", login))

	userStructure := new(user.User)
	err := row.Scan(&userStructure.Name,
		&userStructure.Login,
		&userStructure.PasswordHash,
		&userStructure.CreatedAt,
		&userStructure.LastChangedAt)

	if err == sql.ErrNoRows {
		return nil, internal.ErrNotExistUser
	}
	return userStructure, nil
}

// Извлечение всех логинов
func (u *UserStore) GetAllUsers() ([]*user.User, error) {
	q := `SELECT name, login, password_hash, created_at, change_at
FROM users`
	rows, err := u.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User

	for rows.Next() {
		u := new(user.User)
		err = rows.Scan(
			&u.Name,
			&u.Login,
			&u.PasswordHash,
			&u.CreatedAt,
			&u.LastChangedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// Удаление пользователя по логину
func (u *UserStore) Delete(login string) error {
	q := `DELETE FROM users WHERE login=:loginVal`
	if _, err := u.db.Exec(q,
		sql.Named("loginVal", login)); err != nil {
		return err
	}
	return nil
}

func (u *UserStore) UpdateLogin(login string, newLogin string) error {
	q := `UPDATE users SET login=:newLoginVal WHERE login=:loginVal`
	if _, err := u.db.Exec(q,
		sql.Named("newLoginVal", newLogin),
		sql.Named("loginVal", login)); err != nil {
		return err
	}
	return nil
}

func (u *UserStore) UpdateTime(login string) error {
	q := `UPDATE users SET change_at=:newTime WHERE login=:loginVal`
	if _, err := u.db.Exec(q,
		sql.Named("newTime", time.Now().Format(time.DateTime)),
		sql.Named("loginVal", login)); err != nil {
		return err
	}
	return nil
}

/*
CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL DEFAULT "",
	login VARCHAR(128) NOT NULL DEFAULT "",
	password_hash VARCHAR(128) NOT NULL DEFAULT "",
	created_at VARCHAR(32) NOT NULL DEFAULT "",
	change_at VARCHAR(32) NOT NULL DEFAULT ""
);
*/
