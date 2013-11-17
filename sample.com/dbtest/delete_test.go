package dbtest

import (
	"testing"
)

func TestDelete(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	no := 0
	if row, err := db.Query(`delete from character where no = 5 RETURNING no`); err != nil {
		t.Errorf("Delete Query Error %v", err)
	} else {
		if ok := row.Next(); ok {
			if names, err := row.Columns(); err != nil {
				t.Errorf("row.Columns error %v", err)
			} else {
				for name := range names {
					t.Logf("name = %s", name)
				}
			}
			if err := row.Scan(&no); err != nil {
				t.Errorf("Scan after delete Error %v", err)
			}
		} else {
			t.Logf("Next() = false. No Data?")
		}
		t.Logf("no = %d", no)
	}

}