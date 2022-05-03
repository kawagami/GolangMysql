package mysql

import (
	"database/sql"
	"fmt"
	"mods/setting"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type VideoActresses struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Actress   string         `json:"actress"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type VideoActressesSlice []VideoActresses

func (va *VideoActresses) Get() []VideoActresses {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	tableName := "video_actresses"
	sqlQuery := fmt.Sprintf("select * from %s", tableName)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	var slice []VideoActresses
	for selectQuery.Next() {
		var oneRowData VideoActresses

		err = selectQuery.Scan(
			&oneRowData.Id,
			&oneRowData.Title,
			&oneRowData.Actress,
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

func (va *VideoActresses) Exist(fileTitle string) bool {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	fileTitleString := fileTitle + "%"
	sqlQuery := fmt.Sprintf("SELECT * FROM video_actresses WHERE `title` LIKE '%s'", fileTitleString)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	var slice []VideoActresses
	for selectQuery.Next() {
		var oneRowData VideoActresses

		err = selectQuery.Scan(
			&oneRowData.Id,
			&oneRowData.Title,
			&oneRowData.Actress,
			&oneRowData.CreatedAt,
			&oneRowData.UpdatedAt,
		)
		if err != nil {
			panic(err.Error())
		}
		slice = append(slice, oneRowData)
	}
	//
	return len(slice) > 1
}

func (va *VideoActresses) Insert(data VideoActresses) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	stmt, err := db.Prepare("INSERT INTO video_actresses(title, actress, created_at) VALUES( ?, ?, ? )")
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	if _, err := stmt.Exec(
		// newId,
		data.Title,
		data.Actress,
		time.Now().Format(setting.TIMENOW),
		// data.UpdatedAt,
	); err != nil {
		panic(err)
	}
}

func (va *VideoActresses) Where(column, operator, value string) []VideoActresses {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, err := sql.Open("mysql", sqlOpenString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	sqlQuery := fmt.Sprintf("select * from video_actresses where `%s` %s '%s'", column, operator, value)
	selectQuery, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer selectQuery.Close()
	//
	var slice []VideoActresses
	for selectQuery.Next() {
		var oneRowData VideoActresses

		err = selectQuery.Scan(
			&oneRowData.Id,
			&oneRowData.Title,
			&oneRowData.Actress,
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
func (va *VideoActresses) Delete(column, operator, value string) {
	sqlOpenString := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PW, IP, DB)
	db, _ := sql.Open("mysql", sqlOpenString)
	sqlQuery := fmt.Sprintf("delete from video_actresses where `%s` %s '%s'", column, operator, value)
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
