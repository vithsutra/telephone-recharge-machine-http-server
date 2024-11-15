package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
)

type DataBaseInitUseCase struct {
	dbService *service.DataBaseService
}

func NewDatabaseInitUseCase(dbRepo repository.DataBaseRepository) *DataBaseInitUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &DataBaseInitUseCase{
		dbService,
	}
}

func (u *DataBaseInitUseCase) Execute() (error, int) {
	if err := u.dbService.InitializeDataBase(); err != nil {
		log.Printf("error occurred while initializing the database,database init,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}
	return nil, 0
}
