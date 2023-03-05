package controller

/*
*Controller for the routes
 */
import (
	"log"
	"net/http"

	"github.com/chrishayes042/api/repo"
)

// mappings to the routes
func HandleRequest() {
	http.HandleFunc("/", repo.HomePage)
	http.HandleFunc("/articles", repo.AllArticles)
	http.HandleFunc("/users", repo.GetAllUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
