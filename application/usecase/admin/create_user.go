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
	if request.UserName == "" {
		return fmt.Errorf("user name cannot be empty"), 1
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1
	}

	if err := validation.ValidatePassword(request.Password); err != nil {
		return err, 1
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
		UserId:   userId,
		UserName: request.UserName,
		Password: string(hashedPasswordBytes),
	}

	if err := u.dbService.CreateUser(user); err != nil {
		log.Printf("error occurred with database while creating the user, create user, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
