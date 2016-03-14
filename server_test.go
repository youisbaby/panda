package panda

import (
	"testing"
)

func ControllerTest(c *Controller) {
	ret := struct {
		Name string
		Age  int
	}{}
	c.Write([]byte(RenderForm(&ret)))
}

// go test -run=TestServer
func TestServer(t *testing.T) {
	HandlerRouter(CASUAL, "/.html", ControllerTest)
	Run()
}
