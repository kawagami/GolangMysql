package sqlGorm

import (
	"gorm.io/gorm"
)

// 舊 DB 的 struct
type VideoActress struct {
	gorm.Model
	Title   string
	Actress string
}

// 新 DB 的 table
// 影片番號
type Video struct {
	gorm.Model
	Name      string
	ActressId int
	StorageId int
	// Actresses []Actress `gorm:"many2many:video_actresses;"`
	// Storages  []Storage `gorm:"many2many:video_storages;"`
}

// 演員
type Actress struct {
	gorm.Model
	Name   string
	Videos []Video
}

// 存放位置
type Storage struct {
	gorm.Model
	Name   string
	Videos []Video
}

/*
偷懶不拉關聯 一張表紀錄所有東西 2022-06-18
*/
type VideoMix struct {
	gorm.Model
	Title     string
	LongTitle string
	Actress   string
	CoverImg  string
	RawHtml   string
}
