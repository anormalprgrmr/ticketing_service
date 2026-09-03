package handlers

import "net/http"

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
