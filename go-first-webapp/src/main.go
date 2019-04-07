package main

import (
	"fmt"
	"net/http"
)
var path string="/myfirstgoweb"
func main() {

	http.HandleFunc(path,handle)
	http.ListenAndServe(":8080",nil)
}

func handle(wrt http.ResponseWriter,r *http.Request){
	fmt.Fprintf(wrt,"Hiyya!!!Welcome To First Go Web")
}