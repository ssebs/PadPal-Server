// main.go
package main

import (
	"fmt"
	"log"

	"github.com/ssebs/padpal-server/data"
)

func main() {
	fmt.Println("PadPal Server")

	provider2, err := data.NewFileProvider(`F:\LocalProgramming\_DATA_TEST_\example`)
	if err != nil {
		log.Fatal(err)
	}
	testNote := data.NewNote("Test-Note", "Seb", "# test-note\nGotta love me some test, *amirite*\n")
	err = provider2.SaveNote(testNote)
	if err != nil {
		log.Fatal(err)
	}

	// api.HandleAndServe("", 5000)
}
