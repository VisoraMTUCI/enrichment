package handler

import "net/http"

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!"))
}
