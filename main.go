package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	route struct {
		path string
		view string
	}
)

var routes = []route{
	{
		path: "/",
		view: "views/index.html",
	},
}

func render(w http.ResponseWriter, view string) {
	data, err := ioutil.ReadFile(view)
	if err != nil {
		log.Print("Error", err)
		fmt.Fprint(w, "internal error")
		return
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	path := r.URL.Path
	for _, router := range routes {
		if router.path == path {
			render(w, router.view)
		} else {
			render(w, "views/404.html")
		}
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":1488", nil))
}
