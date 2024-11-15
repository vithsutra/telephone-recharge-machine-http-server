package user

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetMachineExpenseHistoryUseCase struct {
	dbService *service.DataBaseService
}

func NewGetMachineExpenseHistoryUseCase(dbRepo repository.DataBaseRepository) *GetMachineExpenseHistoryUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetMachineExpenseHistoryUseCase{
		dbService,
	}
}

func (u *GetMachineExpenseHistoryUseCase) Execute(machineId string) (error, int, []*entity.ExpenseHistory) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1, nil
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while checking machine id exists,get expense history,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exists"), 1, nil
	}

	expenseHistory, err := u.dbService.GetExpenseHistoryByMachineId(machineId)

	if err != nil {
		log.Printf("error occurred with database while getting machine expense history,get machine expense, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, expenseHistory
}
