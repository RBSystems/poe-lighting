package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/poe-lighting/handlers"
)

func main() {
	log.SetLevel("info")
	port := ":8030"
	router := common.NewRouter()

	// Functionality Endpoints
	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	write.GET("/:address/:light/on", handlers.TurnOn)
	write.GET("/:address/:light/off", handlers.TurnOff)
	write.GET("/:address/:light/:dimlevel", handlers.DimLight)

	// log level endpoints
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
