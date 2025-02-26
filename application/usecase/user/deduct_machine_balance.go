package user

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
)

type DeductMachineBalanceUseCase struct {
	dbService *service.DataBaseService
}

func NewDeductMachineBalanceUseCase(dbRepo repository.DataBaseRepository) *DeductMachineBalanceUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &DeductMachineBalanceUseCase{
		dbService,
	}
}

func (u *DeductMachineBalanceUseCase) Execute(machineId string, request *request.MachineBalanceDeductRequest) (error, int) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1
	}

	if request.Amount == 0 {
		return fmt.Errorf("deduct amount cannot be empty"), 1
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while checking machine id exists, deduct machine balance, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exists"), 1
	}

	balance, err := u.dbService.GetMachineBalance(machineId)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("error occured with database"), 2
	}

	if balance < request.Amount {
		return fmt.Errorf("insufficient funds"), 1
	}

	if err := u.dbService.DeductMachineBalance(machineId, request.Amount); err != nil {
		log.Printf("error occurred with database while deducting the machine balance, deduct machine balance, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
