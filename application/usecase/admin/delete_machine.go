package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type DeleteMachineUseCase struct {
	dbService *service.DataBaseService
}

func NewDeleteMachineUseCase(dbRepo repository.DataBaseRepository) *DeleteMachineUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &DeleteMachineUseCase{
		dbService,
	}
}

func (u *DeleteMachineUseCase) Execute(machineId string) (error, int) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while checking machine id exists, delete machine , Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exists"), 1
	}

	if err := u.dbService.DeleteMachine(machineId); err != nil {
		log.Printf("error occurred with database while deleting the machine, delete machine, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
