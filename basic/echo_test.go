package basic_test

import (
	gc "gopkg.in/check.v1"
	"testing"

	"github.com/wupeka/stork/basic"
)

func Test(t *testing.T) { gc.TestingT(t) }

type EchoSuite struct{}

var _ = gc.Suite(&EchoSuite{})

func (s *EchoSuite) TestEchoOutput(c *gc.C) {
	call := basic.EchoCall{}
	in := basic.EchoInputOutput{Data: "foo"}
	status, out := call.POST(nil, nil, &in)
	c.Check(status, gc.Equals, 200)
	c.Check(out, gc.DeepEquals, in)
}
