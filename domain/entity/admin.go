package entity

type Admin struct {
	AdminId   string `json:"admin_id"`
	AdminName string `json:"admin_name"`
	Password  string `json:"-"`
}
