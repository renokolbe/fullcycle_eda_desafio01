package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	get_accountbalance "github.com/renokolbe/fc-ms-wallet-balance/internal/usecase/get_account-balance"
)

type WebAccountBalanceHandler struct {
	GetAccountBalanceUseCase get_accountbalance.GetAccountBalanceUseCase
}

func NewWebAccountBalanceHandler(getAccountBalanceUseCase get_accountbalance.GetAccountBalanceUseCase) *WebAccountBalanceHandler {
	return &WebAccountBalanceHandler{
		GetAccountBalanceUseCase: getAccountBalanceUseCase,
	}
}

func (h *WebAccountBalanceHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//fmt.Println(`id: `, id)
	dto := get_accountbalance.GetAccountBalanceInputDTO{
		ID: id,
	}
	outputDTO, err := h.GetAccountBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	if outputDTO.Balance == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "account not found"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
