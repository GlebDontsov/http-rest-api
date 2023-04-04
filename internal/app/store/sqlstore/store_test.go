package sqlstore_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=postgres dbname=restapi_test"
		passwordDB, exists := os.LookupEnv("DATABASE_PASSWORD")
		if exists {
			databaseURL = "host=localhost port=5432 user=postgres dbname=restapi_test" + fmt.Sprintf(" password=%s", passwordDB)
		}
	}
	os.Exit(m.Run())
}
