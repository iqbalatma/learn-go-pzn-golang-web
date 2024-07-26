package learn_go_pzn_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello world")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Helo template html")
}
