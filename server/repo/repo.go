package repo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/chrishayes042/api/model"
)

func getEnvVariable(key string) string {
	err := godotenv.Load("postgres.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var host string = getEnvVariable("host")
var port string = getEnvVariable("port")
var user string = getEnvVariable("user")
var pass string = getEnvVariable("pass")
var dbname string = getEnvVariable("db")

func connectToDb() string { //*sql.DB
	psqlInfo := fmt.Sprintf("host=%s, port=%s, user=%s, pass=%s, dbname=%s", host, port, user, pass, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Connected")
	// return db
	return psqlInfo

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
	str := connectToDb()
	fmt.Fprint(w, str)
}
