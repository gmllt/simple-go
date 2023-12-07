package web

import (
	"net/http"
)

func (w *Web) hello(wr http.ResponseWriter, _ *http.Request) {
	WriteJSON(wr, map[string]interface{}{
		"message": "Hello World!",
	})
}
