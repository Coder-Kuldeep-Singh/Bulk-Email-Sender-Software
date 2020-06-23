package main

import (
	"github.com/emails/routers"
)

func main() {
	//setup routes
	r := routers.SetupRouters()
	r.Run(":1000")
}
