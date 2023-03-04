package repo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chrishayes042/model"
)

// get all Articles. Should be a db call
// TODO: create postgresql db
func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := &model.Articles{
		model.Article{Title: "test title", Desc: "test desc", Content: "test content"},
		model.Article{Title: "test title2", Desc: "test desc2", Content: "hello Contente2"},
	}
	fmt.Println("Endpoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)

}

// function to test the homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage endpoint hit")
}
