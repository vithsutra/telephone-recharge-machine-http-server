package user

import (
	"fmt"
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/usecase/utils"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginUseCase struct {
	dbService *service.DataBaseService
}

func NewUserLoginUseCase(dbRepo repository.DataBaseRepository) *UserLoginUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &UserLoginUseCase{
		dbService,
	}
}

func (u *UserLoginUseCase) Execute(request *request.UserLoginRequest) (error, int, string) {
	if request.UserName == "" {
		return fmt.Errorf("user name cannot be empty"), 1, ""
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1, ""
	}

	isUserNameExists, err := u.dbService.CheckUserNameExists(request.UserName)

	if err != nil {
		log.Printf("error occurred with database while checking user name exists, user login, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	if !isUserNameExists {
		return fmt.Errorf("user name not exists"), 1, ""
	}

	user, err := u.dbService.GetUserByUserName(request.UserName)

	if err != nil {
		log.Printf("error occurred with database while getting the user by user nam, user login, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return fmt.Errorf("incorrect password"), 1, ""
	}

	jwtToken, err := utils.CreateJwtToken(user.UserId, user.MachineId, user.UserName)

	log.Println(err)

	if err != nil {
		log.Printf("error occurred while generating the jwt token,admin login,Error -> %v\n", err.Error())
		return fmt.Errorf("error occured while generating token"), 2, ""
	}

	return nil, 0, jwtToken

}
