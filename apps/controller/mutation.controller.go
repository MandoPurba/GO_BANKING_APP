package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MandoPurba/rest-api/apps/service"
	"github.com/MandoPurba/rest-api/utils"
	"net/http"
)

func Transfer(w http.ResponseWriter, r *http.Request) {
	var dto service.MutationDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid mutation payload",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// CEK SALDO
	amount, tax, err := service.Amount(dto.FromAccount)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusInternalServerError,
			Message: "An error occurred while processing the data",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	if amount <= dto.Amount+tax {
		res := utils.Response{
			Code:    http.StatusBadRequest,
			Message: "make sure your balance is enough",
			Data: map[string]interface{}{
				"balance":         amount,
				"tax":             tax,
				"minimum balance": fmt.Sprintf("your balance must be %d for transfers %d", dto.Amount+tax+1, dto.Amount),
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	// TAX DEDUCTION
	dto.Balance = dto.Amount - tax

	_, err = service.Transfer(dto)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to transfer to %s", dto.ToAccount),
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := utils.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"info": fmt.Sprintf("transfer to %s for %d was successful", dto.ToAccount, dto.Amount),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
