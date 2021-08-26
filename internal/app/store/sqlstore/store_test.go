package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain (m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=postgres password=root host=localhost dbname=user_api_test_db sslmode=disable"
	}

	os.Exit(m.Run())
}