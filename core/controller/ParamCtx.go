package controller

import (
	"net/url"
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type ParamCtx struct {
	QueryMap url.Values
	PostJsonMap map[string]interface{}
	PostFormMaP url.Values
}

func parameterFactory(r *http.Request) *ParamCtx {
	res := &ParamCtx{
		QueryMap:nil,
		PostJsonMap:nil,
		PostFormMaP:nil,
	}
	switch r.Method {
    case http.MethodGet:
        {
            res.QueryMap ,_ = url.ParseQuery(r.URL.RawQuery)
		}
	case http.MethodPost:
		{
			contentType := r.Header.Get("Content-Type")
			if contentType == "application/x-www-form-urlencoded" {
				r.ParseForm()
				res.PostFormMaP = r.PostForm
			}
			if contentType == "application/json" {
				res.PostJsonMap = jsonValueHandler(r.Body)
			}

		}
	default:
		{

		}
	}
	return res
}

func jsonValueHandler(r io.ReadCloser) (map[string]interface{}) {
	respBytes, err := ioutil.ReadAll(r)
    if err != nil {
        fmt.Println(err.Error())
        return nil
	}
	var jsonMap map[string]interface{}
	ok := json.Unmarshal(respBytes, &jsonMap)
	if ok != nil {
		fmt.Println(ok.Error())
		return nil
	}
	return jsonMap
}