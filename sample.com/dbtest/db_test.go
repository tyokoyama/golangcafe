package dbtest

import (
	_ "github.com/lib/pq"
	"database/sql"
	"testing"
)

func getConnection() (*sql.DB, error) {
	return sql.Open("postgres", "user=gdgchugoku dbname=sampledb sslmode=verify-full")
}

func TestOpen(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()
}