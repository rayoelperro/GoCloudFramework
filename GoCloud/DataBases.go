package GoCloud

import (
	"database/sql"
	"strings"

	_ "./mysql"
)

type DataBase struct {
	db *sql.DB
}

type DataTable struct {
	db    *sql.DB
	table string
}

func Connect(user string, pass string, dbn string) DataBase {
	db, err := sql.Open("mysql", user+":"+pass+"@/"+dbn+"?charset=utf8")
	checkErr(err)
	return DataBase{db}
}

func (d *DataBase) GetTable(tablename string) DataTable {
	return DataTable{d.db, tablename}
}

func (d *DataTable) Insert(values []string) {
	if len(values) > 0 {
		var q string
		for i := 0; i < len(values); i++ {
			q += "'" + values[i] + "',"
		}
		q = q[:len(q)-1]
		stmt, err := d.db.Prepare("INSERT INTO " + d.table + " VALUES(" + q + ")")
		checkErr(err)

		_, err = stmt.Exec()
		checkErr(err)
	}
}

func (d *DataTable) Update(name string, value string, condition string) {
	stmt, err := d.db.Prepare("update " + d.table + " set " + name + "=? where " + condition)
	checkErr(err)

	_, err = stmt.Exec(value)
	checkErr(err)
}

func (d *DataTable) Query(toget []string, condition string) [][]string {
	rows, err := d.db.Query("SELECT " + strings.Join(toget, ",") + " FROM " + d.table + " where " + condition)
	checkErr(err)
	var result [][]string
	for rows.Next() {
		number, _ := rows.Columns()
		toas := make([]string, len(number))
		row := make([]interface{}, len(number))
		for x := 0; x < len(number); x++ {
			row[x] = &toas[x]
		}
		err = rows.Scan(row...)
		checkErr(err)
		result = append(result, toas)
	}
	return result
}

func (d *DataTable) QueryAll(toget []string) [][]string {
	rows, err := d.db.Query("SELECT " + strings.Join(toget, ",") + " FROM " + d.table)
	checkErr(err)
	var result [][]string
	for rows.Next() {
		number, _ := rows.Columns()
		toas := make([]string, len(number))
		row := make([]interface{}, len(number))
		for x := 0; x < len(number); x++ {
			row[x] = &toas[x]
		}
		err = rows.Scan(row...)
		checkErr(err)
		result = append(result, toas)
	}
	return result
}

func (d *DataTable) Delete(condition string) {
	stmt, err := d.db.Prepare("delete from " + d.table + " where " + condition)
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
