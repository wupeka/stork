package api

import "reflect"

type Call interface {
	// Returns path that is registered for this call
	Path() string
	// Returns an ACL that's minimal to get to this call
	ACL() *Acl
}

type Gettable interface {
	Call
	GET(*Client, map[string]string) (int, interface{})
}

type Puttable interface {
	Call
	InputType() reflect.Type
	PUT(*Client, map[string]string, interface{}) (int, interface{})
}

type Postable interface {
	Call
	InputType() reflect.Type
	POST(*Client, map[string]string, interface{}) (int, interface{})
}
