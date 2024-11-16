package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Router adalah implementasi dari HttpHandler, sehingga bisa digunakan di HttpServer
// Router seperti ServeMux, keleihannya bisa menambahkan method GET, POST, PUT, dll
// httprouter.New() -> untuk membuat router, merupakan pointer, agar selalu menggunakan Objek Router yang sama
// ketika menambahkan route, tidak menggunakan http.handler() seperti di ServeMux, tetapi menggunakan httprouter.handle()
// dimana httprouter.handle() ada tiga parameter -> writer, request, params (untuk menangkap path variabel bukan query param)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}