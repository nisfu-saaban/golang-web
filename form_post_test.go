package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	reqBody := strings.NewReader("first_name=Bambang&last_name=Yuwono")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", reqBody)
	request.Header.Add("content-Type", "application/x-www-form-urlencoded")

	recoder := httptest.NewRecorder()

	FormPost(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
