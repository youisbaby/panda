package panda

import (
	"net/http"
)

//控制器
type Controller struct {
	*http.Request       //请求信息
	http.ResponseWriter //输出信息
}

//创建控制器
func newController(r *http.Request, w http.ResponseWriter) *Controller {
	var c Controller
	c.Request = r
	c.ResponseWriter = w
	return &c
}
func (c *Controller) NotFound() {
	notFound(c.ResponseWriter)
}
func (c *Controller) Write(b []byte) (int, error) {
	return c.ResponseWriter.Write(b)
}
