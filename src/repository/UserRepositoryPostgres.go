package repository

import "fmt"

type UserRepositoryPostgres struct {
	DatabaseConnection string
}

func NewUserRepositoryPostgres(databaseConnection string) UserRepositoryPostgres {
	return UserRepositoryPostgres{}
}

func (u *UserRepositoryPostgres) FindById(id string) string {
	query := `SELECT * FROM users WHERE id = ` + id
	fmt.Println("Query: ", query)
	return "User found"
}
