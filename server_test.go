package panda

import (
	"testing"
)

func ControllerTest(c *Controller) {
	c.Write([]byte("ControllerTest"))
}

// go test -run=TestServer
func TestServer(t *testing.T) {
	HandlerRouter(CASUAL, "/.html", ControllerTest)
	Run()
}
