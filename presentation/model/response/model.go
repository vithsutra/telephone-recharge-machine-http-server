package response

import "github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"

type StatusMessage struct {
	Message string `json:"message"`
}

type AmdinId struct {
	AdminId string `json:"admin_id"`
}

type Users struct {
	Users []*entity.User `json:"users"`
}

type Machines struct {
	Machines []*entity.Machine `json:"machines"`
}

type MachineIds struct {
	MachineIds []string `json:"machine_ids"`
}

type RechargeHistory struct {
	RechargeHistory []*entity.RechargeHistory `json:"recharge_history"`
}

type Token struct {
	Token string `json:"token"`
}

type Balance struct {
	Balance int32 `json:"balance"`
}

type ExpenseHistory struct {
	ExpenseHistory []*entity.ExpenseHistory `json:"expense_history"`
}
