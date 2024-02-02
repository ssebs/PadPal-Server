// main.go
package main

import (
	"fmt"
	"log"

	"github.com/ssebs/padpal-server/data"
)

func main() {
	fmt.Println("PadPal Server")

	provider2, err := data.NewFileProvider("./data/example/")
	if err != nil {
		log.Fatal(err)
	}
	testNote := data.NewNote("Test-Note", "Seb", "# test-note\nGotta love me some test, *amirite*\n")
	provider2.SaveNote(testNote)

	// api.HandleAndServe("", 5000)
}
