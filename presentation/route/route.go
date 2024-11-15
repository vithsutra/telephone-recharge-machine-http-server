package route

import (
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/handler/admin"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/handler/user"
	"github.com/gorilla/mux"
)

func Router(dbRepo repository.DataBaseRepository) *mux.Router {
	adminHandler := admin.NewHandler(dbRepo)
	userHandler := user.NewHandler(dbRepo)

	router := mux.NewRouter()

	rootRouter := router.PathPrefix("/root").Subrouter()
	loginRouter := router.PathPrefix("/login").Subrouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()

	rootRouter.HandleFunc("/database/init", adminHandler.DataBaseInitializeHandler).Methods("GET")
	rootRouter.HandleFunc("/create/admin", adminHandler.CreateAdminHandler).Methods("POST")
	rootRouter.HandleFunc("/delete/admin/{adminId}", adminHandler.DeleteAdminHandler).Methods("GET")

	loginRouter.HandleFunc("/admin", adminHandler.AdminLoginHandler).Methods("POST")
	loginRouter.HandleFunc("/user", userHandler.UserLoginHandler).Methods("POST")

	adminRouter.HandleFunc("/create/user", adminHandler.CreateUserHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/user/{userId}", adminHandler.DeleteUserHandler).Methods("GET")
	adminRouter.HandleFunc("/users", adminHandler.GetAllUsersHandler).Methods("GET")
	adminRouter.HandleFunc("/create/machine/{adminId}", adminHandler.CreatMachineHandler).Methods("POST")
	adminRouter.HandleFunc("/machines/{adminId}", adminHandler.GetMachinesHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/machine/{machineId}", adminHandler.DeleteMachineHandler).Methods("GET")
	adminRouter.HandleFunc("/recharge/machine/{machineId}", adminHandler.RechargeMachineHandler).Methods("POST")
	adminRouter.HandleFunc("/recharge/history/{machineId}", adminHandler.GetRechargeHistoryHandler).Methods("GET")
	adminRouter.HandleFunc("/expense/history/{machineId}", adminHandler.GetExpenseHistoryHandler).Methods("GET")

	userRouter.HandleFunc("/machine/balance/{machineId}", userHandler.GetMachineBalanceHandler).Methods("GET")
	userRouter.HandleFunc("/deduct/machine/balance/{machineId}", userHandler.DeductMachineBalanceHandler).Methods("POST")
	userRouter.HandleFunc("/expense/history/{machineId}", userHandler.GetMachineExpenseHistoryHandler).Methods("GET")

	return router
}
