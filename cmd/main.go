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
	provider2.SaveNote(&data.Note{Title: "test"})

	// api.HandleAndServe("", 5000)
}
