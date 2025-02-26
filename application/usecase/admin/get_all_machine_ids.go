package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type getAllMachineIdsUseCase struct {
	dbService *service.DataBaseService
}

func NewGetAllMachineIdsUseCase(dbRepo repository.DataBaseRepository) *getAllMachineIdsUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &getAllMachineIdsUseCase{
		dbService,
	}
}

func (u *getAllMachineIdsUseCase) Execute(adminId string) (error, int, []string) {

	machineIds, err := u.dbService.GetMachineIdsByAdminId(adminId)

	if err != nil {
		log.Printf("error occurred with database: %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, machineIds
}
