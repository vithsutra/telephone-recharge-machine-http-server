package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetAllUsersUseCase struct {
	dbService *service.DataBaseService
}

func NewGetAllUsersUseCase(dbRepo repository.DataBaseRepository) *GetAllUsersUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetAllUsersUseCase{
		dbService,
	}
}

func (u *GetAllUsersUseCase) Execute() (error, int, []*entity.User) {
	users, err := u.dbService.GetAllUsers()
	if err != nil {
		log.Printf("error occurred with database while getting all the users, get all users, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}
	return nil, 0, users
}
