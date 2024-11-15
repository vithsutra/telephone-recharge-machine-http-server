package user

import (
	"encoding/json"
	"net/http"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/usecase/user"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/request"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/response"
	"github.com/gorilla/mux"
)

type Handler struct {
	dbRepo repository.DataBaseRepository
}

func NewHandler(dbRepo repository.DataBaseRepository) *Handler {
	return &Handler{
		dbRepo,
	}
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var request request.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	userLoginUseCase := user.NewUserLoginUseCase(h.dbRepo)

	err, errStatus, token := userLoginUseCase.Execute(&request)

	if err != nil {
		response := &response.StatusMessage{
			Message: err.Error(),
		}

		if errStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := response.Token{
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetMachineBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]

	getMachineBalanceUseCase := user.NewGetMachineBalanceUseCase(h.dbRepo)

	err, errStatus, balance := getMachineBalanceUseCase.Execute(machineId)

	if err != nil {
		response := &response.StatusMessage{
			Message: err.Error(),
		}

		if errStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.Balance{
		Balance: balance,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeductMachineBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]

	var request request.MachineBalanceDeductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	deductMachineBalanceUseCase := user.NewDeductMachineBalanceUseCase(h.dbRepo)

	err, errStatus := deductMachineBalanceUseCase.Execute(machineId, &request)

	if err != nil {
		response := &response.StatusMessage{
			Message: err.Error(),
		}

		if errStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.StatusMessage{
		Message: "amount deducted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetMachineExpenseHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]

	getExpenseHistoryUseCase := user.NewGetMachineExpenseHistoryUseCase(h.dbRepo)

	err, errStatus, expenseHistory := getExpenseHistoryUseCase.Execute(machineId)

	if err != nil {
		response := &response.StatusMessage{
			Message: err.Error(),
		}

		if errStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := response.ExpenseHistory{
		ExpenseHistory: expenseHistory,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
