package main

import "tublessin/api_gateway/config/router"

func main() {
	muxRouter := router.CreateRouter()
	router.NewAppRouter(muxRouter).InitRouter()
	router.StartServer(muxRouter)
}
