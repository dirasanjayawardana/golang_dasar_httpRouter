package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Router adalah implementasi dari HttpHandler, sehingga bisa digunakan di HttpServer
// Router seperti ServeMux, keleihannya bisa menambahkan method GET, POST, PUT, dll
// httprouter.New() -> untuk membuat router, merupakan pointer, agar selalu menggunakan Objek Router yang sama
// ketika menambahkan route, tidak menggunakan http.handler() seperti di ServeMux, tetapi menggunakan httprouter.handle()
// dimana httprouter.handle() ada tiga parameter -> writer, request, params (untuk menangkap path variabel bukan query param)

func TestRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))

}