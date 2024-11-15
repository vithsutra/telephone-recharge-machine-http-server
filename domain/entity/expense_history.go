package entity

type ExpenseHistory struct {
	MachineId string `json:"machine_id"`
	Amount    int32  `json:"amount"`
	Timestamp string `json:"timestamp"`
}
