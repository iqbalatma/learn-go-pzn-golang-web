package learn_go_pzn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHallo(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Hello World")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=iqbal", nil)
	recorder := httptest.NewRecorder()

	SayHallo(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestMultipleQueryParamValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=iqbal&name=atma", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
