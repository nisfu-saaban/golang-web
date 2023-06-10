package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestSayHello(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/hello?name=nisfu", nil)
	record := httptest.NewRecorder()

	sayHello(record, request)
	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}
