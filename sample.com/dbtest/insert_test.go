package dbtest

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	// 最大Noを取得
	no := 0
	if err := db.QueryRow("select Max(no) from character").Scan(&no); err != nil {
		t.Errorf("QueryRow Error %v", err)
	} else {
		no++
//		if row := db.QueryRow(fmt.Sprintf("insert into character values(%d, '羽柴秀吉', 2.5, 1, 8, 6, '撹乱貫通射撃', 4) RETURNING no", no)); row != nil {
		if row, err := db.Query(fmt.Sprintf("insert into character values(%d, '羽柴秀吉', 2.5, 1, 8, 6, '撹乱貫通射撃', 4) RETURNING no", no)); err != nil {
			t.Errorf("Insert Query Error %v", err)
		} else {
			if names, err := row.Columns(); err != nil {
				t.Errorf("row.Columns error %v", err)
			} else {
				for name := range names {
					t.Logf("name = %s", name)
				}
			}

			row.Next()
			if names, err := row.Columns(); err != nil {
				t.Errorf("row.Columns error %v", err)
			} else {
				for name := range names {
					t.Logf("name = %s", name)
				}
			}

			if err := row.Scan(&no); err != nil {
				t.Errorf("Scan after insert Error %v", err)
			}
			t.Logf("no = %d", no)
		}
	}
}

func TestPrepare(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	// 最大Noを取得
	no := 0
	if err := db.QueryRow("select Max(no) from character").Scan(&no); err != nil {
		t.Errorf("QueryRow Error %v", err)
	} else {
		if st, err := db.Prepare("insert into character values($1, $2, $3, $4, $5, $6, $7, $8)"); err != nil {
			t.Errorf("Prepare Error %v", err)
		} else {
			for i := no + 1; i < (no + 6); i++ {
				if res, err := st.Exec(i, "羽柴秀吉", 2.5, 1, 8, 6, "撹乱貫通射撃", 4); err != nil {
					t.Errorf("Exec Error %v", err)
				} else {
					var lastId int64
					var affectedRows int64
					var err error
					if lastId, err = res.LastInsertId(); err != nil {
						t.Logf("res.LastInsertId() error %v", err)
					}
					if affectedRows, err = res.RowsAffected(); err != nil {
						t.Logf("res.RowsAffected() error %v", err)
					}
					t.Logf("LastInsertId = %d, RowsAffected = %d", lastId, affectedRows)
				}
			}
		}
	}	
}