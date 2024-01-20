package money

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/flpcastro/virtual-wallet/internal/database"
	"github.com/flpcastro/virtual-wallet/internal/user"
	"github.com/google/uuid"
)

type transferRequest struct {
	DebtorID      string `json:"debtor_id"`
	BeneficiaryID string `json:"beneficiary_id"`
	Amount        int    `json:"amount"`
}

func UsersHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	conn, err := database.CreateConnection()
	if err != nil {
		responseFromError(err, w)
		return
	}
	defer conn.Close()

	id := strings.TrimPrefix(req.URL.Path, "/users/")
	userID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	repo := user.Repository{
		Conn: conn,
	}

	balance, err := repo.SelectBalanceByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}
