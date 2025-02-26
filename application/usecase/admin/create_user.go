package admin

import (
	"fmt"
	"log"
	"strings"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/usecase/utils"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	dbService *service.DataBaseService
}

func NewCreateUserUseCase(dbRepo repository.DataBaseRepository) *CreateUserUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &CreateUserUseCase{
		dbService,
	}
}

func (u *CreateUserUseCase) Execute(request *request.CreateUserRequest) (error, int) {

	if request.AdminId == "" {
		return fmt.Errorf("admin id cannot be empty"), 1
	}

	if request.MachineId == "" {
		return fmt.Errorf("machine id cannot be empty"), 1
	}

	if request.Email == "" {
		return fmt.Errorf("email cannot be empty"), 1
	}

	if request.UserName == "" {
		return fmt.Errorf("user name cannot be empty"), 1
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1
	}

	if err := utils.ValidateEmail(request.Email); err != nil {
		return err, 1
	}

	if err := utils.ValidatePassword(request.Password); err != nil {
		return err, 1
	}

	isAdminIdExists, err := u.dbService.CheckAdminIdExists(request.AdminId)

	if err != nil {
		log.Printf("error occurred with database, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isAdminIdExists {
		return fmt.Errorf("admin id not exists"), 1
	}

	isUserNameExists, err := u.dbService.CheckUserNameExists(request.UserName)

	if err != nil {
		log.Printf("error occurred with database while checking user name exists, create user, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if isUserNameExists {
		return fmt.Errorf("user name already exists"), 1
	}

	userId := uuid.New().String()
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 5)

	if err != nil {
		log.Printf("error occurred while generating the hash for password,create user,Error ->%v\n", err.Error())
		return fmt.Errorf("error occurred while generating the password hash"), 2
	}

	user := &entity.User{
		UserId:    userId,
		AdminId:   request.AdminId,
		MachineId: request.MachineId,
		Email:     request.Email,
		UserName:  request.UserName,
		Password:  string(hashedPasswordBytes),
	}

	if err := u.dbService.CreateUser(user); err != nil {
		log.Printf("error occurred with database while creating the user, create user, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if err := utils.SendUserAccessEmail(user.Email, strings.ToUpper(user.MachineId), user.UserName, request.Password); err != nil {
		log.Printf("error occurred while sending email: Error -> %v\n", err)
		return nil, 2
	}

	return nil, 0
}
