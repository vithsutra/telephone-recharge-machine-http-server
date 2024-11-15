package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
)

type RechargeMachineUseCase struct {
	dbService *service.DataBaseService
}

func NewRechargeMachineUseCase(dbRepo repository.DataBaseRepository) *RechargeMachineUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &RechargeMachineUseCase{
		dbService,
	}
}

func (u *RechargeMachineUseCase) Execute(machineId string, request *request.RechargeMachineRequest) (error, int) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1
	}

	if request.Amount == 0 {
		return fmt.Errorf("recharge amount cannot be zero"), 1
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while checking machine id exists, recharge machine, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exists"), 1
	}

	if err := u.dbService.RechargeMachine(machineId, request.Amount); err != nil {
		log.Printf("error occurred with database while recharging the machine, recharge machine, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
