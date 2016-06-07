package gotesting

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goutility"
	"testing"
)

func TestMysql(t *testing.T) {
	fmt.Println("connect to db")

	db, err := sql.Open("mysql", "root:123@/testdb")
	goutility.CheckErr(err)

	rows, err := db.Query("select * from person")

	for rows.Next() {
		var name string
		var age int
		var id string

		err = rows.Scan(&name, &age, &id)
		goutility.CheckErr(err)
		fmt.Println(" name ", " age ", "     id       ")
		fmt.Println(name, age, id)
	}
}
