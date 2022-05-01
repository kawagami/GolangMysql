package mysql

import (
	"database/sql"
	"fmt"
	"mods/setting"
	"time"

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

func (tt *TempTable) Get() []TempTable {
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

func (tt *TempTable) Insert(data TempTable) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	stmt, err := db.Prepare("INSERT INTO temp_table(basename, person, size, file_type, raw_data, location, backed_up, log_time, created_at) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	// // get max id
	// var rawData TempTable
	// allData := rawData.Get()
	// length := len(allData)
	// maxIndex := length - 1
	// newId := allData[maxIndex].Id + 1

	if _, err := stmt.Exec(
		// newId,
		data.Basename,
		data.Person,
		data.Size,
		data.FileType,
		data.RawData,
		data.Location,
		data.BackedUp,
		time.Now().Unix(),
		time.Now().Format(setting.TIMENOW),
		// data.UpdatedAt,
	); err != nil {
		panic(err)
	}
}

func (tt *TempTable) InsertAll(data TempTableSlice) {
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

// func (tt *TempTable) Where(column, operator, value string) []TempTable {
// 	return tt.Get()
// }

func (tt *TempTable) Where(column, operator, value string) []TempTable {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	sqlQuery := fmt.Sprintf("select * from temp_table where `%s` %s '%s'", column, operator, value)
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
func (tt *TempTable) Delete(column, operator, value string) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	sqlQuery := fmt.Sprintf("delete from temp_table where `%s` %s '%s'", column, operator, value)
	result, err := db.Exec(sqlQuery)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	fmt.Println("delete data successd:", result)

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("delete RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Println("delete Affected rows:", rowsaffected)
}
