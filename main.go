package main

import (
	"database/sql"
	"fmt"
	"mods/mysql"
	"time"
)

func main() {
	// var ca mysql.ComicAuthors
	// fmt.Println(len(ca.Get()))

	var tt mysql.TempTable
	fmt.Println(len(tt.Get()))

	slice := mysql.TempTableSlice{}
	newData := mysql.TempTable{
		Id:        len(tt.Get()) + 1,
		Basename:  "新Basename",
		Person:    "新Person",
		Size:      "新Size",
		FileType:  "新FileType",
		RawData:   "新RawData",
		Location:  sql.NullString{},
		BackedUp:  sql.NullBool{},
		LogTime:   int(time.Now().Unix()),
		CreatedAt: sql.NullString{},
		UpdatedAt: sql.NullString{},
	}
	slice = append(slice, newData)
	tt.InsertAll(slice)
}
