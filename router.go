package panda

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	GET    = "GET"  //GET请求
	POST   = "POST" //POST请求
	CASUAL = ""     //任意请求
)

//路由器
type Router struct {
	Method string            //GET|POST请求方式
	Path   string            //请求路径
	Action func(*Controller) //控制器
}

//路由器列表
var routers map[string]*Router = make(map[string]*Router, 0)

//添加路由器
func HandlerRouter(method, url string, f func(*Controller)) {
	_, ok := routers[url]
	if ok {
		//路由已存在，抛出异常
		panic(errors.New(fmt.Sprintf("url:%sexist?", url)))
	}
	if f == nil {
		//函数不能为空，抛出异常
		panic(errors.New("func is nul?"))
	}
	var r Router
	r.Action = f
	r.Path = url
	r.Method = strings.ToUpper(method)
	routers[url] = &r
}
func handle(w http.ResponseWriter, r *http.Request) {
	//匹配URL
	reg := regexp.MustCompile(`\?.*`)
	url := reg.ReplaceAllString(r.RequestURI, "")

	val, ok := routers[url]

	if !ok {
		notFound(w)
		return
	}
	if val.Method != "" && strings.ToUpper(r.Method) != val.Method {
		notFound(w)
		return
	}

	c := newController(r, w)
	//拦截器调用
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			c.Write([]byte(fmt.Sprintf("%s", recoverErr)))
		}
	}()
	interceptorRun(c, val.Action)

}

//路由过滤器
func RouterFilter(c *Controller) {

}
func notFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Fount 404!"))
}
