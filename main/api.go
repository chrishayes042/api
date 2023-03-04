package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chrishayes042/model"
)

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := &model.Articles{
		model.Article{Title: "test title", Desc: "test desc", Content: "test content"},
		model.Article{Title: "test title2", Desc: "test desc2", Content: "hello Contente2"},
	}
	fmt.Println("Endpoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage endpoint hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
