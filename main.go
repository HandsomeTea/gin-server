package main

import (
	"gin-server/server"
	v1 "gin-server/server/routers/v1"
)

func main() {
	r := server.CreateServer()

	v1.RegisterRoutesAll(r)
	r.Run(":8084")
}
