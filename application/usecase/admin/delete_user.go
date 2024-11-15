package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type DeleteUserUseCase struct {
	dbService *service.DataBaseService
}

func NewDeleteUserUseCase(dbRepo repository.DataBaseRepository) *DeleteUserUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &DeleteUserUseCase{
		dbService,
	}
}

func (u *DeleteUserUseCase) Execute(userId string) (error, int) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1
	}

	isUserIdExists, err := u.dbService.CheckUserIdExists(userId)

	if err != nil {
		log.Printf("error occurred with database while checking user id exists,delete user,Error->%v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1
	}

	if err := u.dbService.DeleteUser(userId); err != nil {
		log.Printf("error occurred with database while deleting the user, delete user, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
