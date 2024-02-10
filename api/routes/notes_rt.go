// notes_rt.go - /notes/ routing handlers
package routes

import (
	"fmt"

	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/data/providers"
)

/*
GET
x	/notes?q=
x	/notes/:id
	/notes/:id?version=
	TODO: /versions/notes
POST
x	/notes
PUT
	/notes/:id
	/notes/:id?version=
DELETE
	/notes/:id
	TODO: /notes/:id?version=
*/

// GET //

// GETNotesHandler uses a CRUDProvider and handles GET /notes?q=
// To be used in gin's router.GET()
func GETNotesHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		qry := c.Query("q")

		notes, err := provider.ListNotes(qry)
		if err != nil {
			ErrorHandler(404, err, c)
			return
		}
		c.JSON(200, notes)
	}
}

// GETNoteByIDHandler uses a CRUDProvider and handles GET /notes/:id
// To be used in gin's router.GET()
func GETNoteByIDHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	// Get ID from param
	return func(c *gin.Context) {
		// Parse id as GUID if possible
		id, err := guid.ParseString(c.Param("id"))
		if err != nil {
			ErrorHandler(400, fmt.Errorf("invalid id given, could not convert to guid: err: %s", err), c)
			return
		}
		// Then get the note from that GUID & return
		note, err := provider.LoadNote(*id)
		if err != nil {
			ErrorHandler(404, err, c)
			return
		}
		c.JSON(200, note)
	}
}

// POST //

// POSTNotesHandler uses a CRUDProvider and handles POST /notes
// To be used in gin's router.GET()
func POSTNotesHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Map post data to NoteBind, then create Note from that
		var nb data.NoteBind
		if err := c.ShouldBind(&nb); err != nil {
			ErrorHandler(400, err, c)
			return
		}
		note := data.NewNoteFromBind(nb)

		// Save the new note
		err := provider.SaveNote(note)
		if err != nil {
			ErrorHandler(500, err, c)
			return
		}
		c.JSON(201, note)
	}
}

// PUT //
// To be used in gin's router.PUT()
func PUTNoteHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// DELETE //
// To be used in gin's router.DELETE()
func DELETENoteHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
