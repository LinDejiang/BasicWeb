package gee

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine 路由
type Engine struct {
	router map[string]HandlerFunc
}

// New 获取web实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

//增加路由请求
func (engine *Engine) addRouter(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRouter("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(res, req)
	} else {
		res.Write([]byte("404 not Find"))
	}
}
