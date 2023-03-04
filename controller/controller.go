package controller

import (
	"log"
	"net/http"

	"github.com/chrishayes042/repo"
)

// mappings to the routes
func HandleRequest() {
	http.HandleFunc("/", repo.HomePage)
	http.HandleFunc("/articles", repo.AllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
