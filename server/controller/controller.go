package controller

import (
	"github.com/chrishayes042/api/repo"
	"log"
	"net/http"
)

// mappings to the routes
func HandleRequest() {
	http.HandleFunc("/", repo.HomePage)
	http.HandleFunc("/articles", repo.AllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
