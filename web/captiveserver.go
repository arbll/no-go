package web

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}