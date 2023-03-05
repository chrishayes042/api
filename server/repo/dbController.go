package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load("postgres.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func ConnectToDb() *sqlx.DB { //
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
