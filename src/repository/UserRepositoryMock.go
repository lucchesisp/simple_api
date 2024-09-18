package repository

type UserRepositoryLocal struct {
	DatabaseConnection string
}

func NewUserRepositoryLocal(databaseConnection string) UserRepositoryLocal {
	return UserRepositoryLocal{}
}

func (u *UserRepositoryLocal) FindById(id string) string {
	if id == "123" {
		return ""
	}

	return "Gabriel"
}
