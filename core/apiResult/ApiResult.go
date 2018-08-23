package apiResult

import (
	"encoding/json"
	"fmt"
)

type ApiResult struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type ApiResultInterface interface {
	Success(data map[string]string)
	Error(code int, msg string, data map[string]string)
	Test()
}

const SUCCESS = "SUCCESS"

func New() * ApiResult {
	return &ApiResult{
		Code:0,
		Msg:"",
		Data:nil,
	}
}

func (r * ApiResult) Success(data map[string]interface{}) *ApiResult {
	r.Code = 200
	r.Msg = SUCCESS
	r.Data = data
	return r
}

func (r *ApiResult) Error(code int, msg string,data map[string]interface{}) *ApiResult {
	r.Code = code
	r.Msg = msg
	r.Data = data
	return r
}

func (r *ApiResult) Test() *ApiResult{
	r.Code = 0
	r.Data = nil
	r.Msg = "UNKNOWN"
	return r
}

func (r *ApiResult) JsonBytes() []byte {
	json,err := json.Marshal(r)
	if err != nil{
		fmt.Println("json fail!")
		return nil
	}
	return json

}




