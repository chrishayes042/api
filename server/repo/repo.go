package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chrishayes042/api/model"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnvVariable(key string) string {
	err := godotenv.Load("postgres.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var database *sql.DB = connectToDb()

// var host = getEnvVariable("host")
// var port = getEnvVariable("port")
// var user = getEnvVariable("user")
// var pass = getEnvVariable("pass")
// var dbname = getEnvVariable("db")

func connectToDb() *sql.DB { //
	var (
		host     = getEnvVariable("HOST")
		user     = getEnvVariable("USER")
		password = getEnvVariable("PASS")
		dbname   = getEnvVariable("DBNAME")
		port     = getEnvVariable("PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	// return database connection
	return db

}
func QueryDB(db *sql.DB, query string) *sql.Rows {
	fmt.Println("about to query data")
	rows, err := db.Query(query)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	return rows

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all users")
	rows := QueryDB(database, "select * from users")
	// defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user)
		userArray := &model.Users{}
		if err != nil {
			log.Fatal(err)
		}
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
