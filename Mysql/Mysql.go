package Mysql

import (
	"database/sql"
	"fmt"
)

const (
	USER = "root"
	PW   = "root"
	IP   = "127.0.0.1:3306"
	DB   = "arrange"
)

type ComicAuthors struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

func SelectAllFrom(tableName string) (dataSlice []string) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	sqlQuery := fmt.Sprintf("select * from %s", tableName)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	// var dataSlice []string
	for selectQuery.Next() {
		var comicAuthors ComicAuthors

		err = selectQuery.Scan(&comicAuthors.Id, &comicAuthors.Name, &comicAuthors.CreatedAt, &comicAuthors.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		dataSlice = append(dataSlice, comicAuthors.Name)
		// fmt.Println(comicAuthors.Id)
		// fmt.Println(comicAuthors.Name)
		// fmt.Println(comicAuthors.CreatedAt)
		// fmt.Println(comicAuthors.UpdatedAt)
		// fmt.Println()
	}
	//
	return
}
