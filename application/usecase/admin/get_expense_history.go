package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetExpenseHistoryUseCase struct {
	dbService *service.DataBaseService
}

func NewGetExpenseHistoryUseCase(dbRepo repository.DataBaseRepository) *GetExpenseHistoryUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetExpenseHistoryUseCase{
		dbService,
	}
}

func (u *GetExpenseHistoryUseCase) Execute(machineId string) (error, int, []*entity.ExpenseHistory) {
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
