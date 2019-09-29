package basic

import (
	"github.com/wupeka/stork/api"
	"reflect"
)

type EchoInputOutput struct {
	Data string `json:data`
}

type EchoCall struct{}

func (c *EchoCall) Path() string {
	return "/echo"
}

func (c *EchoCall) ACL() *api.Acl {
	return nil
}

func (c *EchoCall) InputType() reflect.Type {
	return reflect.TypeOf(EchoInputOutput{})
}

func (c *EchoCall) POST(client *api.Client, _ map[string]string, r interface{}) (int, interface{}) {
	in := r.(*EchoInputOutput)
	return 200, EchoInputOutput{in.Data}
}

var _ api.Call = (*EchoCall)(nil)
var _ api.Postable = (*EchoCall)(nil)
