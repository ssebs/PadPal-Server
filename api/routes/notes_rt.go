// notes_rt.go - /notes/ routing handlers
package routes

import (
	"fmt"
	"time"

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

		notes, err := provider.LoadNotes(qry)
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
		// Validate incoming data
		if err := validateNoteBind(nb); err != nil {
			c.JSON(400, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// Create note from the incoming data
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

// validateNoteBind will validate the incoming data from an HTTP request
// returns an error
// TODO: support combining multiple errors
func validateNoteBind(nb data.NoteBind) error {
	if nb.Author == "" {
		return fmt.Errorf("%s is missing the author field", nb)
	}
	if nb.Title == "" {
		return fmt.Errorf("%s is missing the title field", nb)
	}

	return nil
}

// PUT //
// To be used in gin's router.PUT()
func PUTNoteHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse id as GUID if possible
		id, err := guid.ParseString(c.Param("id"))
		if err != nil {
			ErrorHandler(400, fmt.Errorf("invalid id given, could not convert to guid: err: %s", err), c)
			return
		}
		// Then get the note from that GUID & return
		n, err := provider.LoadNote(*id)
		if err != nil {
			ErrorHandler(404, err, c)
			return
		}
		// copy the data
		note := *n

		// Get what fields were updated and update them

		// Map post data to NoteBind, then create Note from that
		var nb data.NoteBind
		if err := c.ShouldBind(&nb); err != nil {
			ErrorHandler(400, err, c)
			return
		}

		// Compare data & update if needed
		if nb.Title != "" && note.Title != nb.Title {
			note.Title = nb.Title
			// TODO: if title is changed, the file get's duplicated..
			// TODO: delete the old file
		}
		if nb.Author != "" && note.Author != nb.Author {
			note.Author = nb.Author
		}
		if nb.Contents != "" && note.Contents != nb.Contents {
			note.Contents = nb.Contents
		}
		note.LastUpdated = time.Now().UTC()

		// Update on the FS
		// del note
		if err = provider.DeleteNote(*note.ID); err != nil {
			ErrorHandler(500, err, c)
		}
		// save new one
		err = provider.UpdateNote(*note.ID, &note)
		if err != nil {
			ErrorHandler(400, err, c)
		}

		c.JSON(201, note)
	}
}

// DELETE //
// To be used in gin's router.DELETE()
func DELETENoteHandler(provider providers.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
