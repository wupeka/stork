package basic

import (
	"github.com/wupeka/stork/api"
)

type VersionOutput struct {
	Version string `json:version`
	BuildID string `json:buildid`
}

type VersionCall struct{}

func (c *VersionCall) Path() string {
	return "/version"
}

func (c *VersionCall) ACL() *api.Acl {
	return nil
}

func (c *VersionCall) GET(client *api.Client, _ map[string]string) (int, interface{}) {
	return 200, VersionOutput{Version: "1.2.3", BuildID: "512341234"}
}

var _ api.Gettable = (*VersionCall)(nil)
