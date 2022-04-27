package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TempTable struct {
	Id        int            `json:"id"`
	Basename  string         `json:"basename"`
	Person    string         `json:"person"`
	Size      string         `json:"size"`
	FileType  string         `json:"file_type"`
	RawData   string         `json:"raw_data"`
	Location  sql.NullString `json:"location"`
	BackedUp  sql.NullBool   `json:"backed_up"`
	LogTime   int            `json:"log_time"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type TempTableSlice []TempTable

func (this *TempTable) Get() []TempTable {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	tableName := "temp_table"
	sqlQuery := fmt.Sprintf("select * from %s", tableName)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	var slice []TempTable
	for selectQuery.Next() {
		var oneRowData TempTable

		err = selectQuery.Scan(
			&oneRowData.Id,
			&oneRowData.Basename,
			&oneRowData.Person,
			&oneRowData.Size,
			&oneRowData.FileType,
			&oneRowData.RawData,
			&oneRowData.Location,
			&oneRowData.BackedUp,
			&oneRowData.LogTime,
			&oneRowData.CreatedAt,
			&oneRowData.UpdatedAt,
		)
		if err != nil {
			panic(err.Error())
		}
		slice = append(slice, oneRowData)
	}
	//
	return slice
}

func (this *TempTable) InsertAll(data TempTableSlice) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	stmt, err := db.Prepare("INSERT INTO temp_table(id, basename, person, size, file_type, raw_data, location, backed_up, log_time, created_at, updated_at) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, tt := range data {
		if _, err := stmt.Exec(
			tt.Id,
			tt.Basename,
			tt.Person,
			tt.Size,
			tt.FileType,
			tt.RawData,
			tt.Location,
			tt.BackedUp,
			tt.LogTime,
			tt.CreatedAt,
			tt.UpdatedAt,
		); err != nil {
			panic(err)
		}
	}
}
