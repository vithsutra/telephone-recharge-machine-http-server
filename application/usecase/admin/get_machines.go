package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetMachinesUseCase struct {
	dbService *service.DataBaseService
}

func NewGetMachinesUseCase(dbRepo repository.DataBaseRepository) *GetMachinesUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetMachinesUseCase{
		dbService,
	}
}

func (u *GetMachinesUseCase) Execute(adminId string) (error, int, []*entity.Machine) {
	if adminId == "" {
		return fmt.Errorf("admin id cannot be empty"), 1, nil
	}

	isAdminIdExists, err := u.dbService.CheckAdminIdExists(adminId)

	if err != nil {
		log.Printf("error occurred with database while checking admin id exists, get machines, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isAdminIdExists {
		return fmt.Errorf("admin id not exists"), 1, nil
	}

	machines, err := u.dbService.GetMachinesByAdminId(adminId)

	if err != nil {
		log.Printf("error occurred with database while getting machines by admin id, get machines, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, machines
}
