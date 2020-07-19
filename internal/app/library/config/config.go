package config

import (
	"app/library/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var RainDog *database.Config

func init()  {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, ip, port, dbName)
	dsn = dsn + "?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"

	RainDog = &database.Config{
		Addr:         "rain_dog",
		DSN:          dsn,
		Active:       10,
		Idle:         5,
		IdleTimeout:  time.Duration(time.Minute),
		QueryTimeout: time.Duration(time.Minute),
		ExecTimeout:  time.Duration(time.Minute),
		TranTimeout:  time.Duration(time.Minute),
	}

}

