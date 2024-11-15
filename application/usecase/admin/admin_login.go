package admin

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/service"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginUseCase struct {
	dbService *service.DataBaseService
}

func NewAdminLoginUseCase(dbRepo repository.DataBaseRepository) *AdminLoginUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &AdminLoginUseCase{
		dbService,
	}
}

func (u *AdminLoginUseCase) Execute(request *request.AdminLoginRequest) (error, int, string, string) {
	if request.AdminName == "" {
		return fmt.Errorf("admin name cannot be empty"), 1, "", ""
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1, "", ""
	}

	isAdminNameExists, err := u.dbService.CheckAdminNameExists(request.AdminName)

	if err != nil {
		log.Printf("error occurred with database while checking admin name exists,admin login,Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, "", ""
	}

	if !isAdminNameExists {
		return fmt.Errorf("admin name not exists"), 1, "", ""
	}

	admin, err := u.dbService.GetAdminByAdminName(request.AdminName)

	if err != nil {
		log.Printf("error occurred with database while getting admin, admin login, Error -> %v\n", err.Error())
		return fmt.Errorf("error occurred with database"), 2, "", ""
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password)); err != nil {
		return fmt.Errorf("incorrect password"), 1, "", ""
	}

	secreteKey := os.Getenv("SECRETE_KEY")

	if secreteKey == "" {
		log.Printf("missing or empty env variable SECRETE_KEY\n")
		return fmt.Errorf("secrete key not found"), 2, "", ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_name": request.AdminName,
		"exp":        time.Now().Add(time.Hour * 24 * 360).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secreteKey))

	if err != nil {
		log.Printf("error occurred while generating the jwt token,admin login,Error -> %v\n", err.Error())
		return fmt.Errorf("error occured while generating token"), 2, "", ""
	}

	return nil, 0, tokenString, admin.AdminId
}
