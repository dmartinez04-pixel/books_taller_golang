package routes

import (
	"net/http"
	"server/cmd/handlers"
)

func SetupPingRoute() {
	http.HandleFunc("/ping", handlers.PingHandler)
}
