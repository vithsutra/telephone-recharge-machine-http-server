package admin

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/validation"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateAdminUseCase struct {
	dbService *service.DataBaseService
}

func NewCreateAdminUseCase(dbRepo repository.DataBaseRepository) *CreateAdminUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &CreateAdminUseCase{
		dbService,
	}
}

func (u *CreateAdminUseCase) Execute(request *request.CreateAdminRequest) (error, int) {
	if request.AdminName == "" {
		return fmt.Errorf("admin name cannot be empty"), 1
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1
	}

	if err := validation.ValidatePassword(request.Password); err != nil {
		return err, 1
	}

	isAdminNameExists, err := u.dbService.CheckAdminNameExists(request.AdminName)

	if err != nil {
		log.Printf("error occurred with database while checking admin name exists,create admin,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if isAdminNameExists {
		return fmt.Errorf("admin name already exists"), 1
	}

	adminId := uuid.New().String()
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 5)

	if err != nil {
		log.Printf("error occurred while generating hash password, create admin, Error -> %v\n", err.Error())
		return fmt.Errorf("error occured while generating password hash"), 2
	}

	admin := &entity.Admin{
		AdminId:   adminId,
		AdminName: request.AdminName,
		Password:  string(hashedPasswordBytes),
	}

	if err := u.dbService.CreateAdmin(admin); err != nil {
		log.Printf("error occurred with database while creating the admin,create admin,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
