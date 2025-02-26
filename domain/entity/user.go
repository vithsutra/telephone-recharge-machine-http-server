package entity

type User struct {
	UserId    string `json:"user_id"`
	AdminId   string `json:"-"`
	MachineId string `json:"machine_id"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	Password  string `json:"-"`
}
