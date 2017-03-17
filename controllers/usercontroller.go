package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"
)

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
	context := NewContext()
	defer context.Close()
	cx := context.GetCollection("users")
	repo := &data.UserRepository{cx}


	if err:=repo.CreateUser(user);err!=nil{
		common.DisplayAppError(
			w,
			err,
			"Register fail",
			http.StatusInternalServerError,
		)
	}

	user.HashPassword=nil
	if j,err:=json.Marshal(UserResource{Data:*user});err!=nil{
		common.DisplayAppError(
			w,
			err,
			"An unexcepted error has occurred",
			http.StatusInternalServerError,
		)
	}else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}


func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string

	err:=json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			http.StatusInternalServerError,
		)
		return
	}
	loginModel:=dataResource.Data
	loginUser:=models.User{
		Email:loginModel.Email,
		Password:loginModel.Password,
	}
	context:=NewContext()
	defer context.Close()
	cx:=context.GetCollection("users")
	repo:=&data.UserRepository{cx}

	if user,err:=repo.Login(loginUser);err!=nil{
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	}else {
		token,err=common.GenerateJWT(user.Email,"member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type","application/json")
		user.HashPassword=nil
		authUser:=AuthUserModel{
			User:user,
			Token:token,
		}
		j,err:=json.Marshal(AuthUserResource{Data:authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				http.StatusInternalServerError,
			)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}
