package main 

import (
	"net/http"
	"net/url"
	"fmt"
	// "strings"
	// "odserver"
	// "log"
	"./SimpleServer"
	"./core/apiResult"
	"./core/controller"
)

// func HelloServer1(w http.ResponseWriter, req *http.Request) {

//     fmt.Fprint(w,"hello world")
// }
// func main() {
//     http.HandleFunc("/test", HelloServer1)
//     err := http.ListenAndServe(":23456", nil)
//     if err != nil {
//         log.Fatal("ListenAndServe: ", err.Error())
//     }
// }

func HelloServer1(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w,"hello world")
}

func queryHandler(w http.ResponseWriter, req *http.Request) {
	QueryMap ,_ := url.ParseQuery(req.URL.RawQuery)
	fmt.Println(QueryMap)
    fmt.Fprint(w,QueryMap)
}

func TestJson(w http.ResponseWriter, req *http.Request) {
	r := (apiResult.New()).Success(nil)
	w.Header().Set("Content-type", "application/json")
 	w.Write(r.JsonBytes())
}

func main() {

	uc := new(controller.UserController)
	pc := new(controller.PageController)
	server := SimpleServer.Create()
	server.GET("/test", HelloServer1)
	server.GET("/query", queryHandler)
	server.GET("/test/query", pc.TestQuery)
	server.POST("/json", TestJson)
	server.POST("/user/login", uc.Login)
	server.POST("/user/add", uc.Add)
	server.PORT(":23457")

}
	
