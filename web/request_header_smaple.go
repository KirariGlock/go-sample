package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// curl -id "first_name=sausheong" 127.0.0.1:8080/body みたいなコマンドで呼び出せるよ
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
