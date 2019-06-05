package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "test:x219erqqgmxxlZ@tcp(rm-wz91670r7o0zi042j8o.mysql.rds.aliyuncs.com:3306)/axd_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	stmtOut, err := db.Prepare("select NAME from axd_city WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var cityName string
	err = stmtOut.QueryRow(110100).Scan(&cityName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("the city name of 110100 is: %s", cityName)

}