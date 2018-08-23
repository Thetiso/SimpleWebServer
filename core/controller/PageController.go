package controller

import (
	"net/http"
)

type PageController struct {}

type PageControllerInterface interface {
	Index(HandlerFunc)
	Error(HandlerFunc)
	TestQuery(HandlerFunc)
}

const INDEX_PAGE = "This Is Index Page!"
const ERROR_PAGE = "This Is Error Page!"

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func (ctl *PageController) Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(INDEX_PAGE))
}

func (ctl *PageController) TestQuery(w http.ResponseWriter, req *http.Request) {
	paramCtx := parameterFactory(req)
	newQueryMap := paramCtx.QueryMap
	p := newQueryMap.Get("p")
	w.Write([]byte(p))
}

func (ctl *PageController) Error(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(ERROR_PAGE))
}