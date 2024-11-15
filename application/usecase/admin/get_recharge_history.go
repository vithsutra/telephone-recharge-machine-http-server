package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type GetRechargeHistoryUseCase struct {
	dbService *service.DataBaseService
}

func NewGetRechargeHistoryUseCase(dbRepo repository.DataBaseRepository) *GetRechargeHistoryUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetRechargeHistoryUseCase{
		dbService,
	}
}

func (u *GetRechargeHistoryUseCase) Execute(machineId string) (error, int, []*entity.RechargeHistory) {
	if machineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1, nil
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(machineId)

	if err != nil {
		log.Printf("error occurred with database while checking machine id exists, get recharge history, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isMachineIdExists {
		return fmt.Errorf("machine id not exists"), 1, nil
	}

	rechargeHistory, err := u.dbService.GetRechargeHistoryByMachineId(machineId)

	if err != nil {
		log.Printf("error occurred with database while getting machine recharge history, get recharge history, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, rechargeHistory
}
