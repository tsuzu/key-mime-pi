package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tsuzu/key-mime-pi/app/handler"
	"github.com/tsuzu/key-mime-pi/app/pkg/hid"
	"github.com/tsuzu/key-mime-pi/app/templates"
)

func main() {
	khid, err := hid.NewKeyboardHID("/dev/hidg0")

	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(templates.FS, khid)

	port := "8000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Println("Listening on port :" + port)

	http.ListenAndServe(":"+port, h)
}
