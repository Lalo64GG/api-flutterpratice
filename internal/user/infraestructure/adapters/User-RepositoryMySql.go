










package adapters

import (
	"api/cmd/db"
	"api/internal/user/domain/entities"
	"database/sql"
	"log"
)

type UserRepositoryMySql struct {
	DB *sql.DB
}

func NewUserRepositoryMySql() (*UserRepositoryMySql, error) {
	db, err := db.Connect();
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	return &UserRepositoryMySql{DB: db}, nil
}

func (r *UserRepositoryMySql) Create(user entities.User) (entities.User, error) {
	query := "INSERT INTO users (name, last_name, email, password) VALUES (?, ?, ?, ?)"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	result, err := r.DB.Exec(query, user.Name, user.LastName, user.Email, user.Password)
	if err != nil {
		return entities.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, err
	}

	user.ID = int(id)
	user.Password = ""
	return user, nil
}

func (r *UserRepositoryMySql) GetById(id int64) (entities.User, error){
	query := "SELECT id, name, last_name, email FROM users WHERE id = ?"

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	
	var user entities.User

	err = row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}


func (r *UserRepositoryMySql) GetByEmail(email string) (entities.User, error) {
	query := "SELECT id, name, last_name, email, password FROM users WHERE email = ?"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user entities.User

	err = row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}