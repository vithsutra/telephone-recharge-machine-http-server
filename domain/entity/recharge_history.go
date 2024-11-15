package entity

type RechargeHistory struct {
	MachineId string `json:"-"`
	Amount    int32  `json:"amount"`
	Timestamp string `json:"timestamp"`
}
