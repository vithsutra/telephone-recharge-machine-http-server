package user

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetMachineBalanceUseCase struct {
	dbService *service.DataBaseService
}

func NewGetMachineBalanceUseCase(dbRepo repository.DataBaseRepository) *GetMachineBalanceUseCase {
	dbService := service.NewDataBaseService(dbRepo)

	return &GetMachineBalanceUseCase{
		dbService,
	}
}

func (u *GetMachineBalanceUseCase) Execute(machineId string) (error, int, int32) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1, 0
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while chekcing machine id exists, get machine balance, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, 0
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exsists"), 1, 0
	}

	balance, err := u.dbService.GetMachineBalance(machineId)

	if err != nil {
		log.Printf("error occurred with database while getting the machine balance, get machine balance, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, 0
	}

	return nil, 0, balance
}
