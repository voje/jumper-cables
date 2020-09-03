package main

import (
	"net/http"
	"fmt"
)

var unhealthy bool = false

func root(w http.ResponseWriter, req *http.Request) {
	if !unhealthy{
		fmt.Fprintf(w, `
		<p>Don't click on this button!</p>
		<form action="/kill">
			<input type="submit" value="Kill" />
		</form>
		`)
	}
}

func kill(w http.ResponseWriter, req *http.Request) {
	unhealthy = true
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/kill", kill)

	http.ListenAndServe(":8000", nil)
}