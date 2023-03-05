package repo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chrishayes042/api/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB = ConnectToDb()

// using sqlx to connect to the database

// function to retrieve all users from the users table
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userArray := model.User{}
	rows, _ := database.Queryx("select * from users")

	for rows.Next() {
		err := rows.StructScan(&userArray)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userArray)
		json.NewEncoder(w).Encode(userArray)
	}
}

// get all Articles. Should be a db call
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
	fmt.Fprint(w)
}
