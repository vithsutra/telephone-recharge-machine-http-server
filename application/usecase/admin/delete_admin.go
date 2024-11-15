package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type DeleteAdminUseCase struct {
	dbService *service.DataBaseService
}

func NewDeleteAdminUseCase(dbRepo repository.DataBaseRepository) *DeleteAdminUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &DeleteAdminUseCase{
		dbService,
	}
}

func (u *DeleteAdminUseCase) Execute(adminId string) (error, int) {
	if adminId == "" {
		return fmt.Errorf("admin id cannot be empty"), 1
	}

	isAdminIdExists, err := u.dbService.CheckAdminIdExists(adminId)

	if err != nil {
		log.Printf("error occurred with database while checking admin id exists,delete admin,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isAdminIdExists {
		return fmt.Errorf("admin id not exists"), 1
	}

	if err := u.dbService.DeleteAdmin(adminId); err != nil {
		log.Printf("error occurred while deleting the admin, delete admin, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0

}
