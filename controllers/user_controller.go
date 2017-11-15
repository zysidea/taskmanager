package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"
)

//POST /user/register
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			http.StatusInternalServerError,
		)
		return
	}
	user := &dataResource.Data
	mc := NewMongoContext()
	defer mc.Close()
	context := mc.GetCollection("users")
	repo := &data.UserRepository{context}

	if err := repo.CreateUser(user); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Register fail",
			http.StatusInternalServerError,
		)
		return
	}

	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}

//POST /user/login
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			http.StatusInternalServerError,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	mc := NewMongoContext()
	defer mc.Close()
	context := mc.GetCollection("users")
	repo := &data.UserRepository{context}

	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	} else {
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		if j, err := json.Marshal(AuthUserResource{Data: authUser});err != nil{

			common.DisplayUnexpectedAppError(w, err)
		}else {
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}

}
