package repository

import "github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"

type DataBaseRepository interface {
	CheckAdminNameExists(adminName string) (bool, error)
	CreateAdmin(admin *entity.Admin) error
	GetAdminByAdminName(adminName string) (*entity.Admin, error)
	CheckAdminIdExists(adminId string) (bool, error)
	DeleteAdmin(adminId string) error
	CheckUserNameExists(userName string) (bool, error)
	CreateUser(user *entity.User) error
	CheckUserIdExists(userId string) (bool, error)
	DeleteUser(userId string) error
	GetUserByUserName(userName string) (*entity.User, error)
	GetAllUsers(adminId string) ([]*entity.User, error)
	CheckMachineIdExists(machineId string) (bool, error)
	CreateMachine(machine *entity.Machine) error
	DeleteMachine(machineId string) error
	GetMachinesByAdminId(adminId string) ([]*entity.Machine, error)
	GetMachineIdsByAdminId(adminId string) ([]string, error)
	RechargeMachine(machineId string, amount int32) error
	GetRechargeHistoryByMachineId(machineId string) ([]*entity.RechargeHistory, error)
	GetMachineBalance(machineId string) (int32, error)
	DeductMachineBalance(machineId string, amount int32) error
	GetExpenseHistoryByMachineId(machineId string) ([]*entity.ExpenseHistory, error)
}
