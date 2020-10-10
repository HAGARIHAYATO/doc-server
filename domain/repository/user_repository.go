package repository

import (
	"database/sql"
	"doc-server/domain/model"
)

type userRepository struct{
	conn *sql.DB
}

type UserOption struct {
	Email string
	ID int64
	Limit int
	Offset int
}

type UserRepository interface {
	Fetch(options *UserOption) ([]*model.User, error)
	Create(user *model.User) (*model.User, error)
	FetchByEmail(email string) (*model.User, error)
	FetchByID(id int64) (*model.User, error)
}

func NewUserRepository(Conn *sql.DB) UserRepository {
	return &userRepository{Conn}
}

func (r userRepository) Fetch(options *UserOption) ([]*model.User, error) {
	var users []*model.User
	rows, err := r.conn.Query("SELECT id, name, email FROM users;")
	if rows == nil { return nil, err }
	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err == nil {
			users = append(users, user)
		}
	}
	return users, err
}

func (r *userRepository) FetchByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.conn.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1;", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (r *userRepository) FetchByID(id int64) (*model.User, error) {
	user := &model.User{}
	rows := r.conn.QueryRow("SELECT id, name, email FROM users WHERE id = $1;", id)
	err := rows.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	stmt, err := r.conn.Prepare("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ID int64
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&ID)
	user.ID = ID
	return user, err
}