package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Magowtham/telephone_recharge_machine_http_server/application/usecase/utils"
	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/model/response"
	"github.com/gorilla/mux"
)

type tokenValidationMiddleware struct {
	dbRepo repository.DataBaseRepository
}

func NewTokenValidationMiddleware(dbRepo repository.DataBaseRepository) *tokenValidationMiddleware {
	return &tokenValidationMiddleware{
		dbRepo,
	}
}

func (m *tokenValidationMiddleware) UserTokenValidater(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				response := &response.StatusMessage{
					Message: "missing the token in header",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			arr := strings.Split(authHeader, " ")

			if len(arr) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				response := &response.StatusMessage{
					Message: "invalid header format",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			token := arr[1]

			claims, err := utils.DecodeJwtToken(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				response := &response.StatusMessage{
					Message: "invalid token",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			tokenMachineId, _ := claims["machine_id"].(string)

			requestedMachineId := mux.Vars(r)["machineId"]

			if tokenMachineId != requestedMachineId {
				w.WriteHeader(http.StatusUnauthorized)
				response := &response.StatusMessage{
					Message: "machine id mismatch occurred",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			userId, _ := claims["user_id"].(string)

			isUserExists, err := m.dbRepo.CheckUserIdExists(userId)

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				response := &response.StatusMessage{
					Message: "error occurred with database",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			if !isUserExists {
				w.WriteHeader(http.StatusUnauthorized)
				response := &response.StatusMessage{
					Message: "invalid user",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
