package main

import (
	// "database/sql"
	"fmt"
	"mods/mysql"
	_ "mods/structs"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// data := mysql.SelectAllFrom()
	var ca mysql.ComicAuthors

	for _, v := range ca.Get() {
		fmt.Println(v)
	}
}
