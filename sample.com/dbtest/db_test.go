package dbtest

import (
	_ "github.com/lib/pq"
	"database/sql"
	"testing"
)

func getConnection() (*sql.DB, error) {
//	return sql.Open("postgres", "user=gdgchugoku dbname=sampledb sslmode=verify-full")
	return sql.Open("postgres", "user=gdgchugoku dbname=sampledb sslmode=disable")
}

func TestOpen(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	db.Close()
}

func TestQuery(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	if rows, err := db.Query("select * from character"); err == nil {
		for rows.Next() {
			// 行の中身を見る。
			var no int
			var name string
			var cost float64
			var typeId int
			var attack int
			var lead int
			var scheme string
			var morale int

			if err := rows.Scan(&no, &name, &cost, &typeId, &attack, &lead, &scheme, &morale); err != nil {
				t.Errorf("Scan Error %v", err)
			}
			t.Logf("Rows[%d, %s, %f, %d, %d, %d, %s, %d]", no, name, cost, typeId, attack, lead, scheme, morale)
		}
	} else if err == sql.ErrNoRows {
		t.Logf("Query NoRows")
	} else {
		t.Errorf("Query error %v", err)
	}

}

func TestQueryRow(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	if row := db.QueryRow("select * from character"); row != nil {
		// 行の中身を見る。
		var no int
		var name string
		var cost float64
		var typeId int
		var attack int
		var lead int
		var scheme string
		var morale int

		if err := row.Scan(&no, &name, &cost, &typeId, &attack, &lead, &scheme, &morale); err == nil {
			t.Logf("Rows[%d, %s, %f, %d, %d, %d, %s, %d]", no, name, cost, typeId, attack, lead, scheme, morale)
		} else if err == sql.ErrNoRows {
			t.Logf("QueryRow NoRows")
		} else {
			t.Errorf("Scan error %v", err)
		}

	} else {
		t.Errorf("QueryRow error %v", err)
	}
}

type Sample struct {
	No int
	Name string
	Cost float64
	TypeId int
	Attack int
	Lead int
	Scheme string
	Morale int
}

func TestQueryRowParamStruct(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		t.Errorf("database open error %v", err)
	}
	defer db.Close()

	if row := db.QueryRow("select * from character"); row != nil {
		var st Sample
		// 構造体などを直接指定しても行を読み取る事は不可能。
		if err := row.Scan(&st.No, &st.Name, &st.Cost, &st.TypeId, &st.Attack, &st.Lead, &st.Scheme, &st.Morale); err == nil {
			t.Logf("Rows[%d, %s, %f, %d, %d, %d, %s, %d]", st.No, st.Name, st.Cost, st.TypeId, st.Attack, st.Lead, st.Scheme, st.Morale)
		} else if err == sql.ErrNoRows {
			t.Logf("QueryRow NoRows")
		} else {
			t.Errorf("Scan error %v", err)
		}
	}
}