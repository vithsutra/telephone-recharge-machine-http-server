package service

import (
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
)

type DataBaseService struct {
	repo repository.DataBaseRepository
}

func NewDataBaseService(repo repository.DataBaseRepository) *DataBaseService {
	return &DataBaseService{
		repo,
	}
}

func (s *DataBaseService) InitializeDataBase() error {
	return s.repo.Init()
}

func (s *DataBaseService) CheckAdminNameExists(adminName string) (bool, error) {
	return s.repo.CheckAdminNameExists(adminName)
}

func (s *DataBaseService) CreateAdmin(admin *entity.Admin) error {
	return s.repo.CreateAdmin(admin)
}

func (s *DataBaseService) DeleteAdmin(adminId string) error {
	return s.repo.DeleteAdmin(adminId)
}

func (s *DataBaseService) GetAdminByAdminName(adminName string) (*entity.Admin, error) {
	return s.repo.GetAdminByAdminName(adminName)
}

func (s *DataBaseService) CheckAdminIdExists(adminId string) (bool, error) {
	return s.repo.CheckAdminIdExists(adminId)
}

func (s *DataBaseService) CheckUserNameExists(userName string) (bool, error) {
	return s.repo.CheckUserNameExists(userName)
}

func (s *DataBaseService) CreateUser(user *entity.User) error {
	return s.repo.CreateUser(user)
}

func (s *DataBaseService) DeleteUser(userId string) error {
	return s.repo.DeleteUser(userId)
}

func (s *DataBaseService) CheckUserIdExists(userId string) (bool, error) {
	return s.repo.CheckUserIdExists(userId)
}

func (s *DataBaseService) GetUserByUserName(userName string) (*entity.User, error) {
	return s.repo.GetUserByUserName(userName)
}

func (s *DataBaseService) GetAllUsers() ([]*entity.User, error) {
	return s.repo.GetAllUsers()
}

func (s *DataBaseService) CheckMachineIdExists(machineId string) (bool, error) {
	return s.repo.CheckMachineIdExists(machineId)
}

func (s *DataBaseService) CreateMachine(machine *entity.Machine) error {
	return s.repo.CreateMachine(machine)
}

func (s *DataBaseService) DeleteMachine(machineId string) error {
	return s.repo.DeleteMachine(machineId)
}

func (s *DataBaseService) GetMachinesByAdminId(adminId string) ([]*entity.Machine, error) {
	return s.repo.GetMachinesByAdminId(adminId)
}

func (s *DataBaseService) RechargeMachine(machineId string, amount int32) error {
	return s.repo.RechargeMachine(machineId, amount)
}

func (s *DataBaseService) GetRechargeHistoryByMachineId(machineId string) ([]*entity.RechargeHistory, error) {
	return s.repo.GetRechargeHistoryByMachineId(machineId)
}

func (s *DataBaseService) GetMachineBalance(machineId string) (int32, error) {
	return s.repo.GetMachineBalance(machineId)
}

func (s *DataBaseService) DeductMachineBalance(machineId string, amount int32) error {
	return s.repo.DeductMachineBalance(machineId, amount)
}

func (s *DataBaseService) GetExpenseHistoryByMachineId(machineId string) ([]*entity.ExpenseHistory, error) {
	return s.repo.GetExpenseHistoryByMachineId(machineId)
}
