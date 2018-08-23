package SimpleServer

import (
    "net/http"
    "fmt"
)

type SimpleServer struct {
    router MethodMaps
}


type ISimpleServer interface {
    PORT(addr string)
    GET(url string, f HandlerFunc)
    POST(url string, f HandlerFunc)
    PUT(url string, f HandlerFunc)
    DELETE(url string, f HandlerFunc)
}

type HandlerMapped struct {
    f HandlerFunc
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

const PAGE_NOT_FOUND = "Page Not Found!"

func Create() *SimpleServer {
    return &SimpleServer{
        router:Router(),
    }
}


func (o *SimpleServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

    o.doHandler(w,req)
}

func (o *SimpleServer) doHandler(w http.ResponseWriter, req *http.Request) {    
    handled := false;

    switch req.Method {
    case http.MethodGet:
        {
            //匹配带url参数的路径
            if hm, ok := o.router.GetMapping(req.URL.Path); ok {
                hm.f(w, req)
                handled = true
            }
        }
    case http.MethodPost:
        {
            if hm, ok := o.router.PostMapping(req.URL.RequestURI()); ok {
                hm.f(w, req)
                handled = true
            }

        }
    case http.MethodDelete:
        {
            if hm, ok := o.router.DeleteMapping(req.URL.String()); ok {
                hm.f(w, req)
                handled = true
            }
        }
    case http.MethodPut:
        {
            if hm, ok := o.router.PutMapping(req.URL.String()); ok {
                hm.f(w, req)
                handled = true
            }
        }
    default:
        {
            
        }
    }
    if !handled {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(PAGE_NOT_FOUND))
    }
}

func (o *SimpleServer) PORT(addr string) {
    fmt.Println("start serving on " + addr)
    err := http.ListenAndServe(addr, o)
	if err != nil {
		panic(err)
	}
}

func (o *SimpleServer) GET(url string, f HandlerFunc) {
    o.router.GetAdd(url, HandlerMapped{f: f})
}
func (o *SimpleServer) POST(url string, f HandlerFunc) {
    o.router.PostAdd(url, HandlerMapped{f: f})
}
func (o *SimpleServer) PUT(url string, f HandlerFunc) {
    o.router.PutAdd(url, HandlerMapped{f: f})
}
func (o *SimpleServer) DELETE(url string, f HandlerFunc) {
    o.router.DeleteAdd(url, HandlerMapped{f: f})
}

