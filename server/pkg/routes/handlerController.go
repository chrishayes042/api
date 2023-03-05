package routes

/*
*Controller for the routes
 */
import (
	"log"
	"net/http"

	con "github.com/chrishayes042/api/pkg/controller"
)

// mappings to the routes
func HandleRequest() {
	http.HandleFunc("/", con.HomePage)
	http.HandleFunc("/articles", con.AllArticles)
	http.HandleFunc("/users", con.GetAllUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
