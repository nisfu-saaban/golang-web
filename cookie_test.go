package golangweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-AWG-Name"
	cookie.Value = r.URL.Query().Get("name")

	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-AWG-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Bambang", nil)
	recorder := httptest.NewRecorder()
	SetCookie(recorder, req)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("cookie %s %s\n", cookie.Name, cookie.Value)
	}
}
