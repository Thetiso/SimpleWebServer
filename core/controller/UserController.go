package controller

import (
	"net/http"
	"fmt"
	"../apiResult"
)

type UserController struct {}

type IUserController interface {
	Login(HandlerFunc)
	Add(HandlerFunc)
	Test(req *http.Request)
}

func (uc *UserController) Login(w http.ResponseWriter, req *http.Request) {
	paramCtx := parameterFactory(req)
	userDTO := paramCtx.PostJsonMap
	fmt.Println(userDTO)
	r := (apiResult.New()).Success(userDTO)
	w.Header().Set("Content-type", "application/json")
	w.Write(r.JsonBytes())
}

func (uc *UserController) Add(w http.ResponseWriter, req *http.Request) {
	r := (apiResult.New()).Error(0, "Test Add User!", nil)
	w.Header().Set("Content-type", "application/json")
	w.Write(r.JsonBytes())
}



