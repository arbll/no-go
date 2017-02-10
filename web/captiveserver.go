package web

import (
	"net/http"
	"github.com/spf13/viper"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":" + viper.GetString("webserver.port"), nil)
}