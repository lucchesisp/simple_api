package repository

import "fmt"

type UserRepositoryMongo struct {
	DatabaseConnection string
}

func NewUserRepositoryMongo(databaseConnection string) UserRepositoryMongo {
	return UserRepositoryMongo{}
}

func (u *UserRepositoryMongo) FindById(id string) string {
	query := `MONGOQUERY XUYZ` + id

	fmt.Println("Query: ", query)
	return "User found"
}
