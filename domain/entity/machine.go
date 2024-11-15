package entity

type Machine struct {
	AdminId         string `json:"-"`
	MachineId       string `json:"machine_id"`
	Label           string `json:"label"`
	Balance         int32  `json:"balance"`
	UpdateTimestamp string `json:"update_timestamp"`
}
