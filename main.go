package main

import (
	"go-mongodb/config"
	"go-mongodb/routes"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/rs/cors"
)

func main() {
	port := config.Getenv("PORT")
	color.Cyan("🌏 Server running on localhost:" + port)
	config.InitDB()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	router := routes.Routes()
	handler := c.Handler(router)
	http.ListenAndServe(":"+port, config.LogRequest(handler))
}
