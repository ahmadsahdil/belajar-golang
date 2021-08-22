package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")

	}
	server := http.Server{
		Addr:    "localhost:8383",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}

func TestTermux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/coba", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello Coba")
	})
	server := http.Server{
		Addr:    "localhost:8383",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}
