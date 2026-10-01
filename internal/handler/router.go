package handler

import (
	"fmt"
	"net/http"
	"time"
)

func MuxFactory(handlers map[string]http.HandlerFunc) *http.ServeMux {
	mux := http.NewServeMux()
	for pattern, handler := range handlers {
		mux.HandleFunc(pattern, handler)
	}

	return mux
}

func ServerFactory(
	port int,
	handlers map[string]http.HandlerFunc,
) *http.Server {
	mux := MuxFactory(handlers)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return server
}
