package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/usecase/admin"
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

func (h *Handler) CreateAdminHandler(w http.ResponseWriter, r *http.Request) {
	var request request.CreateAdminRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	createAdminUseCase := admin.NewCreateAdminUseCase(h.dbRepo)

	err, errStatus := createAdminUseCase.Execute(&request)

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
		Message: "admin created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteAdminHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["adminId"]

	deleteAdminUseCase := admin.NewDeleteAdminUseCase(h.dbRepo)

	err, errStatus := deleteAdminUseCase.Execute(adminId)

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
		Message: "admin deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var request request.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	createUserUseCase := admin.NewCreateUserUseCase(h.dbRepo)

	err, errStatus := createUserUseCase.Execute(&request)

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
		Message: "user created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	deleteUserUseCase := admin.NewDeleteUserUseCase(h.dbRepo)

	err, errStatus := deleteUserUseCase.Execute(userId)

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
		Message: "user deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	adminId := mux.Vars(r)["adminId"]

	getAllUsersUseCase := admin.NewGetAllUsersUseCase(h.dbRepo)

	err, errStatus, users := getAllUsersUseCase.Execute(adminId)

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

	response := &response.Users{
		Users: users,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	var request request.AdminLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	adminLoginUseCase := admin.NewAdminLoginUseCase(h.dbRepo)

	err, errStatus, token, adminId := adminLoginUseCase.Execute(&request)

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

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 365),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	response := &response.AmdinId{
		AdminId: adminId,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreatMachineHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	adminId := vars["adminId"]

	var request request.CreateMachineRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	machineCreateUseCase := admin.NewCreateMachineUseCase(h.dbRepo)

	err, errStatus := machineCreateUseCase.Execute(adminId, &request)

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
		Message: "machine created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetMachinesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["adminId"]
	getMachinesUseCase := admin.NewGetMachinesUseCase(h.dbRepo)

	err, errStatus, machines := getMachinesUseCase.Execute(adminId)

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

	response := response.Machines{
		Machines: machines,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]
	deleteMachineUseCase := admin.NewDeleteMachineUseCase(h.dbRepo)
	err, errStatus := deleteMachineUseCase.Execute(machineId)
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

	response := response.StatusMessage{
		Message: "machine deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetMachineIdsHandler(w http.ResponseWriter, r *http.Request) {
	adminId := mux.Vars(r)["adminId"]

	getAllMachineIdsUseCase := admin.NewGetAllMachineIdsUseCase(h.dbRepo)

	err, errStatus, machineIds := getAllMachineIdsUseCase.Execute(adminId)

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

	response := &response.MachineIds{
		MachineIds: machineIds,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (h *Handler) RechargeMachineHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	machineId := vars["machineId"]

	var request request.RechargeMachineRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	rechargeMachineUseCase := admin.NewRechargeMachineUseCase(h.dbRepo)
	err, errStatus := rechargeMachineUseCase.Execute(machineId, &request)
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
		Message: "recharge successfull",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetRechargeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]

	getRechargeHistoryUseCase := admin.NewGetRechargeHistoryUseCase(h.dbRepo)

	err, errStatus, rechargeHistory := getRechargeHistoryUseCase.Execute(machineId)

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

	response := &response.RechargeHistory{
		RechargeHistory: rechargeHistory,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetExpenseHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineId := vars["machineId"]

	getExpenseHistoryUseCase := admin.NewGetExpenseHistoryUseCase(h.dbRepo)

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
