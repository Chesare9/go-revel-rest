package controllers

import (
	"github.com/revel/revel"
	"go-revel-rest/app/models"
	"go-revel-rest/app/repository"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) GetUserById(id string) revel.Result {

	user, err := repository.GetUserRepository().GetUserById(id)

	response := JsonResponse{}
	response.Success = err == nil
	response.Data = user
	if err != nil {
		response.Error = err.Error()
	}

	return c.RenderJson(response)
}

func (c UserController) GetUsers() revel.Result {

	users := repository.GetUserRepository().GetUsers()

	response := JsonResponse{}
	response.Data = users

	return c.RenderJson(response)
}

func (c UserController) SaveUser(id, nickname string) revel.Result {

	user := &models.User{
		Id:       id,
		Nickname: nickname,
	}

	err := repository.GetUserRepository().SaveUser(user)

	response := JsonResponse{}
	response.Success = err == nil
	if err != nil {
		response.Error = err.Error()
	}
	return c.RenderJson(response)
}
