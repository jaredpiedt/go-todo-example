package mysql

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func newTestStore() *Store {
	// Connect to the database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	s, err := NewStore(db)
	if err != nil {
		panic(err)
	}
	return s
}

func randomString(n int) string {
	s := make([]byte, n)
	rand.Seed(int64(time.Now().Nanosecond()))
	rand.Read(s)
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(s)
}

func TestPrepareStatements(t *testing.T) {
	// Connect to the database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}

	unprepared := map[string]string{
		"test": "SELECT 1",
	}

	stmts, err := prepareStmts(db, unprepared)
	if err != nil {
		t.Fatal(err)
	}

	if len(stmts) != 1 {
		t.Fatalf("incorrect number of statements prepared; got %v want %v\n", len(stmts), 1)
	}
}
