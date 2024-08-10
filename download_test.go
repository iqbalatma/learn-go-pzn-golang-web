package learn_go_pzn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")

	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad request")
		return
	}

	writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	http.ServeFile(writer, request, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
