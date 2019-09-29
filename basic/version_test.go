package basic_test

import (
	"github.com/wupeka/stork/basic"
	gc "gopkg.in/check.v1"
)

type VersionSuite struct{}

var _ = gc.Suite(&VersionSuite{})

func (s *VersionSuite) TestVersionOutput(c *gc.C) {
	call := basic.VersionCall{}
	status, out := call.GET(nil, nil)
	c.Check(status, gc.Equals, 200)
	c.Check(out, gc.DeepEquals, basic.VersionOutput{Version: "1.2.3", BuildID: "512341234"})
}
