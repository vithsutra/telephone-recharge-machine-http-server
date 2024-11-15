package admin

import (
	"fmt"
	"log"
	"time"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
)

type CreateMachineUseCase struct {
	dbService *service.DataBaseService
}

func NewCreateMachineUseCase(dbRepo repository.DataBaseRepository) *CreateMachineUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &CreateMachineUseCase{
		dbService,
	}
}

func (u *CreateMachineUseCase) Execute(adminId string, request *request.CreateMachineRequest) (error, int) {
	if request.MachineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1
	}

	if request.Label == "" {
		return fmt.Errorf("label cannot be empty"), 1
	}

	isAdminIdExists, err := u.dbService.CheckAdminIdExists(adminId)

	if err != nil {
		log.Printf("error occurred with database while checking admin id exists, create machine, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isAdminIdExists {
		return fmt.Errorf("admin id not exists"), 1
	}

	isMachineIdExists, err := u.dbService.CheckMachineIdExists(request.MachineId)

	if err != nil {
		log.Printf("error occurred with database while checking user id exists,create machine,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if isMachineIdExists {
		return fmt.Errorf("machine id already exists"), 1
	}

	machine := &entity.Machine{
		AdminId:         adminId,
		MachineId:       request.MachineId,
		Label:           request.Label,
		Balance:         0,
		UpdateTimestamp: time.Now().String(),
	}

	if err := u.dbService.CreateMachine(machine); err != nil {
		log.Printf("error occurred while creating the machine, create machine, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
