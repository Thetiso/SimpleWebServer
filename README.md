我的简书介绍：https://www.jianshu.com/p/b3579eb9bca5

## 起步
```
func HelloServer1(w http.ResponseWriter, req *http.Request) {

    fmt.Fprint(w,"hello world")
}
func main() {
    http.HandleFunc("/test", HelloServer1)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
```
上述代码简单实现了一个httpserver，访问本地127.0.0.1:8080/test返回`hello world`，但是get、post等都返回一致结果，无法限制访问方式。

## router
router的基本功能是需要限制访问路径和访问方式。
访问方式一般有如下几种
- GET
- POST
- DELETE
- PUT
...

仔细查看上述2行代码:
`http.HandleFunc("/test", HelloServer1)`是为一个path指定了处理函数，`http.ListenAndServe(":8080", nil)`则是实现了监听端口，其第二个参数实际为Handler
```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

...

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
so，router的核心就是根据path和method来指向一个预定义的handlerFunc，用map/[ ]来存储和查找就再适合不过了。
```
//router
const (
    GET         = iota
    POST
    PUT
    DELETE
    CONNECTIBNG
    HEAD
    OPTIONS
    PATCH
    TRACE
)

func NewRouter() MethodMaps {
    return []handler{
        GET:  make(handler),
        POST: make(handler),
        PUT: make(handler),
        DELETE: make(handler),
    }
}

type MethodMaps [] handler
type handler map[string]HandlerMapped
...
//存取函数.....
...

//简单server
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
//路径处理相关函数：路径匹配、路径增加
...

```

此时，main函数的内部代码块就有了一种node.js中koa-router的意味
```
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
```
对比下koa中的结构：
![koa-router](https://upload-images.jianshu.io/upload_images/701398-141f5c210a36fdd8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
索性就把controller(未包装的handler)用类似koa中结构进行构造
```
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
```
PageController是基于GET请求查找的页面handler,不可避免需要考虑带参数的情况。URL中的参数可以由以下代码获得：
```
QueryMap ,_ := url.ParseQuery(req.URL.RawQuery)
```
既然要想办法获取get的参数，平时使用最多的莫过于post传递json参数，获取方式就比较麻烦了：
```
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
```
入参的io.ReadCloser其实是http.Request.Body。
将上述两种获取参数的方法整合起来，就有了ParamCtx这个结构，分别单独获取url中的参数、post方法的json参数等
```
type ParamCtx struct {
    QueryMap url.Values
    PostJsonMap map[string]interface{}
    PostFormMaP url.Values
}
```
写到这，又有一个不可避免的问题：`json形式的返回数据`,结构如下
```
{
    "code": 200,
    "msg": "SUCCESS",
    "data": null
}
```
那么就来定义一个新的结构，ApiResult，以及一些常用的方法
```
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
```
最后呢，针对不同的传参进行测试：
![返回json数据](https://upload-images.jianshu.io/upload_images/701398-142c8ccbc4cf2821.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![返回json数据](https://upload-images.jianshu.io/upload_images/701398-b28423fe642c3662.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![url带参数测试](https://upload-images.jianshu.io/upload_images/701398-79ec605d4ab5e681.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)













