package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MandoPurba/rest-api/apps/service"
	"github.com/MandoPurba/rest-api/config"
	"github.com/MandoPurba/rest-api/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetALlUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetALlUsers()
	if err != nil {
		res := utils.Response{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong!",
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
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := service.GetUserById(id)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("User with id %d not found", id),
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := utils.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	p := &config.Param{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	var dto service.UserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid users payload",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	hash, err := config.EncryptPassword(dto.Password, p)
	if err != nil {
		log.Fatal(err)
	}

	dto.Hash = hash

	user, err := service.CreateUser(dto)
	if err != nil {
		res := utils.Response{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong!",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := utils.Response{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
