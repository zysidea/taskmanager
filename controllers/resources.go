package controllers

import "taskmanager/models"

type (
	//POST /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//POST /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//用户登陆之后的response
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	//登陆验证
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//授权用户token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)
