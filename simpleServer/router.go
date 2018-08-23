package SimpleServer

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

func Router() MethodMaps {
    return []handler{
        GET:    make(handler),
        POST:   make(handler),
        PUT:    make(handler),
        DELETE: make(handler),
    }
}

type MethodMaps [] handler
type handler map[string]HandlerMapped

func (m MethodMaps) GetMapping(url string) (HandlerMapped, bool) {
    if hm, ok := m[GET][url]; ok {
        return hm, true
    }
    return HandlerMapped{}, false
}

func (m MethodMaps) PostMapping(url string) (HandlerMapped, bool) {
    if hm, ok := m[POST][url]; ok {
        return hm, true
    }
    return HandlerMapped{}, false
}

func (m MethodMaps) DeleteMapping(url string) (HandlerMapped, bool) {
    if hm, ok := m[DELETE][url]; ok {
        return hm, true
    }
    return HandlerMapped{}, false
}

func (m MethodMaps) PutMapping(url string) (HandlerMapped, bool) {
    if hm, ok := m[PUT][url]; ok {
        return hm, true
    }
    return HandlerMapped{}, false
}

func (m MethodMaps) GetAdd(url string, mapped HandlerMapped) {
    if _, ok := m.GetMapping(url); ok {
        panic("duplicate url with get method")
    }
    m[GET].SetUrl(url,mapped)
}

func (m MethodMaps) PostAdd(url string, mapped HandlerMapped) {
    if _, ok := m.GetMapping(url); ok {
        panic("duplicate url with Post method")
    }
    m[POST].SetUrl(url,mapped)

}

func (m MethodMaps) PutAdd(url string, mapped HandlerMapped) {
    if _, ok := m.GetMapping(url); ok {
        panic("duplicate url with Put method")
    }
    m[PUT].SetUrl(url,mapped)

}

func (m MethodMaps) DeleteAdd(url string, mapped HandlerMapped) {
    if _, ok := m.GetMapping(url); ok {
        panic("duplicate url with Delete method")
    }
    m[DELETE].SetUrl(url,mapped)
}
func (h handler) SetUrl(url string, mapped HandlerMapped) {
    h[url] = mapped
}
