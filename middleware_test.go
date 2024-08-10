package learn_go_pzn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(writer, "hallo middleware")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
