package controllers

import (
	"encoding/json"
	"financify/pkg/models"
	"financify/pkg/utils"
	"net/http"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)

	err := CreateUser.BeforeSave()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	savedUser, err := CreateUser.SaveUser(uc.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, _ := json.Marshal(savedUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email, emailExists := requestData["email"]
	password, passwordExists := requestData["password"]

	if !emailExists || !passwordExists {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	user := &models.User{}
	user, err = user.FindUserByEmail(uc.DB, email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.VerifyPassword(password)

	if err != nil {
		http.Error(w, "Email or password incorrect!", http.StatusUnauthorized)
		return
	}

    tokenString, _:= utils.GenerateToken(uint32(user.ID))

	mapToken := map[string]interface{}{"authToken": tokenString, "user": user}
	res, _ := json.Marshal(mapToken)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (uc *UserController) UpdateAvatar(w http.ResponseWriter, r *http.Request) {

	uid, err := utils.ExtractUserID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData map[string]string

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
        http.Error(w, "Error aca: " + err.Error(), http.StatusBadRequest)
		return
	}

	avatar, avatarExists := requestData["avatar"]

	if !avatarExists {
		http.Error(w, "avatar must be provided", http.StatusBadRequest)
	}

	user := models.User{}

	result, err := user.UpdateUserAvatar(uc.DB, uint32(uid), avatar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
