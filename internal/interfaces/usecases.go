package interfaces

import "github.com/DKhorkov/hmtm-bff/internal/entities"

type UseCases interface {
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
	RegisterUser(userData entities.RegisterUserDTO) (int, error)
	LoginUser(userData entities.LoginUserDTO) (string, error)
}
