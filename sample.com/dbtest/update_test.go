package dbtest

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	no := 0
	if row, err := db.Query(`update character set name='更新しました。' where no = 3 RETURNING no`); err != nil {
		t.Errorf("Update Query Error %v", err)
	} else {
		row.Next()
		if names, err := row.Columns(); err != nil {
			t.Errorf("row.Columns error %v", err)
		} else {
			for name := range names {
				t.Logf("name = %s", name)
			}
		}
		if err := row.Scan(&no); err != nil {
			t.Errorf("Scan after update Error %v", err)
		}
		t.Logf("no = %d", no)
	}

	if row, err := db.Query(`update character set name='複数更新' where no >= 21 and no <=29 RETURNING no`); err != nil {
		t.Errorf("Update Query Error %v", err)
	} else {
		for row.Next() {
			if names, err := row.Columns(); err != nil {
				t.Errorf("row.Columns error %v", err)
			} else {
				for name := range names {
					t.Logf("name = %s", name)
				}
			}
			if err := row.Scan(&no); err != nil {
				t.Errorf("Scan after update Error %v", err)
			}
			t.Logf("no = %d", no)
		}
	}

}
