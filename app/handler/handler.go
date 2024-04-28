package handler

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/tsuzu/key-mime-pi/app/pkg/hid"
)

type Handler struct {
	handler     http.Handler
	keyboardHID hid.KeyboardHID
}

func NewHandler(fs fs.FS, keyboardHID hid.KeyboardHID) *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		handler:     mux,
		keyboardHID: keyboardHID,
	}

	mux.Handle("/", http.FileServer(http.FS(fs)))
	mux.HandleFunc("/ws", h.WebSocketHandler)

	return h
}

func (h *Handler) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if op == ws.OpClose {
				return
			}
			if !op.IsData() {
				return
			}

			var event hid.JEKeyEvent
			err = json.Unmarshal(msg, &event)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			log.Printf("event: %+v", event)

			err = h.writeKeyEvent(event)

			if err != nil {
				log.Printf("failed to write key event: %v", err)
				// continue processing
			}

			msg, err = json.Marshal(map[string]any{
				"success": err == nil,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}()
}

func (h *Handler) writeKeyEvent(event hid.JEKeyEvent) error {
	controlChars, hidKeyCode, ok := hid.Convert(event)
	if !ok {
		return fmt.Errorf("invalid key event: %v", event)
	}

	return h.keyboardHID.WriteKeyboardReport(controlChars, hidKeyCode)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
