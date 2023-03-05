package repo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chrishayes042/api/model"
	"github.com/jmoiron/sqlx"
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

var database *sqlx.DB = connectToDb()

// using sqlx to connect to the database
func connectToDb() *sqlx.DB { //
	var (
		host     = getEnvVariable("HOST")
		user     = getEnvVariable("USER")
		password = getEnvVariable("PASS")
		dbname   = getEnvVariable("DBNAME")
		port     = getEnvVariable("PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to DB %s\n", dbname)
	// return database connection
	return db

}

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
