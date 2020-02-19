package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"



func main() {
	fmt.Println("hello world")
	db, err := sql.Open("mysql", "root:wubo123456@/test123")
	rows, err := db.Query("select * from orders")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	columns, err := rows.Columns()
	fmt.Println(columns)

	valus := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(valus))
	for i := range valus {
		scanArgs[i] = &valus[i]
		fmt.Println(valus[i])
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}

		for _, col := range valus {
			fmt.Println(string(col))
		}

	}
}
