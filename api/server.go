package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wupeka/stork/state"
	"log"
	"reflect"
)

type APIServer struct {
	ge *gin.Engine
	st *state.State
}

func (as *APIServer) makeGET(gettable Gettable) func(*gin.Context) {
	return func(ctx *gin.Context) {
		c := Client{as.st}
		status, response := gettable.GET(&c, nil)
		ctx.JSON(status, response)
	}
}

func (as *APIServer) makePOST(postable Postable) func(*gin.Context) {
	return func(ctx *gin.Context) {
		c := Client{as.st}
		pt := postable.InputType()
		var input interface{}
		if pt != nil {
			d := json.NewDecoder(ctx.Request.Body)
			d.DisallowUnknownFields()
			v := reflect.New(pt)
			input = v.Interface()
			err := d.Decode(input)
			if err != nil {
				log.Print(err)
			}
		}
		status, response := postable.POST(&c, nil, input)
		ctx.JSON(status, response)
	}
}
func (as *APIServer) RegisterCall(call interface{}) error {
	ac := call.(Call)
	path := ac.Path()
	// GetAcl, authorized, etc.
	get, ok := call.(Gettable)
	if ok && get != nil {
		as.ge.GET(path, as.makeGET(get))
	}
	post, ok := call.(Postable)
	if ok && post != nil {
		as.ge.POST(path, as.makePOST(post))
	}
	return nil
}

func (as *APIServer) Run() {
	as.ge.Run()
}

func MakeAPIServer(st *state.State) (APIServer, error) {
	var as APIServer
	as.ge = gin.Default()
	as.st = st
	return as, nil
}
