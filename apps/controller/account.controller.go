package controller

import (
	"encoding/json"
	"github.com/MandoPurba/rest-api/apps/service"
	"github.com/MandoPurba/rest-api/utils"
	"math/rand"
	"net/http"
	"strconv"
)

func ActivateAccount(w http.ResponseWriter, r *http.Request) {
	var dto service.AccountDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid account payload",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	accountNumber := generateAccountNumber(dto.AccountType, dto.Currency, dto.UserId)
	_, err = service.CreateAccount(accountNumber, dto.AccountType, dto.Currency, dto.UserId)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusInternalServerError,
			Message: "failed to activate the account",
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
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func generateAccountNumber(accountType, currency, userId int) string {
	// Generate random 8-digit number
	randomNumber := rand.Intn(90000000) + 10000000

	// Convert integers to strings
	accountTypeStr := strconv.Itoa(accountType)
	randomNumberStr := strconv.Itoa(randomNumber)
	currencyStr := strconv.Itoa(currency)
	userIdStr := strconv.Itoa(userId)

	// Concatenate the parts to form the account number
	accountNumber := accountTypeStr + randomNumberStr + currencyStr + userIdStr

	return accountNumber
}
