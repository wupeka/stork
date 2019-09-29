package main

import (
	"github.com/wupeka/stork/api"
	"github.com/wupeka/stork/basic"
	"github.com/wupeka/stork/state"
	"log"
)

func main() {
	var st state.State
	err := st.Config.Load("stork.yaml")
	if err != nil {
		log.Fatal(err)
	}
	as, err := api.MakeAPIServer(&st)
	if err != nil {
		log.Fatal(err)
	}
	as.RegisterCall(&basic.VersionCall{})
	as.RegisterCall(&basic.EchoCall{})
	as.Run()
}
