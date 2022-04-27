package mysql

import (
	"database/sql"
	"fmt"
)

type ComicAuthors struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

func (this *ComicAuthors) Get() []ComicAuthors {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	tableName := "comic_authors"
	sqlQuery := fmt.Sprintf("select * from %s", tableName)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	var slice []ComicAuthors
	for selectQuery.Next() {
		var comicAuthors ComicAuthors

		err = selectQuery.Scan(&comicAuthors.Id, &comicAuthors.Name, &comicAuthors.CreatedAt, &comicAuthors.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		slice = append(slice, comicAuthors)
		// fmt.Println(comicAuthors.Id)
		// fmt.Println(comicAuthors.Name)
		// fmt.Println(comicAuthors.CreatedAt)
		// fmt.Println(comicAuthors.UpdatedAt)
		// fmt.Println()
	}
	//
	return slice
}
