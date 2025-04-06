package request

type CreateAdminRequest struct {
	AdminName string `json:"admin_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserRequest struct {
	AdminId   string `json:"admin_id"`
	MachineId string `json:"machine_id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type AdminLoginRequest struct {
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
}

type CreateMachineRequest struct {
	MachineId string `json:"machine_id"`
	Label     string `json:"label"`
}

type RechargeMachineRequest struct {
	Amount int32 `json:"amount"`
}

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type MachineBalanceDeductRequest struct {
	Amount int32 `json:"amount"`
}
