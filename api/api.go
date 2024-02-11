// api.go - REST API for PadPal-Server
package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ssebs/padpal-server/api/routes"
	"github.com/ssebs/padpal-server/data/providers"
)

// HandleAndServe will handle the routes and serve HTTP
// contains a list of route handlers
func HandleAndServe(host string, port int, provider providers.CRUDProvider) {
	hostPort := fmt.Sprintf("%s:%d", host, port)

	// TODO: stop using gin
	// init gin + HTTP handlers
	router := gin.Default()
	initHandlers(router, provider)

	// Run the server
	log.Fatal(router.Run(hostPort))
}

// initHandlers is where to define new HTTP handlers
// requires the current router, and a CRUDProvider
func initHandlers(router *gin.Engine, provider providers.CRUDProvider) {
	// Default handler
	router.GET("/", routes.RootHandler)
	// Notes handlers
	router.GET("/notes", routes.GETNotesHandler(provider))
	router.GET("/notes/:id", routes.GETNoteByIDHandler(provider))
	router.PUT("/notes/:id", routes.PUTNoteHandler(provider))
	router.POST("/notes", routes.POSTNotesHandler(provider))

}
