package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/example/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
