package crawler

import (
	"gorm.io/gorm"
)

/*
紀錄從網站抓的 1. 未整理文字 slice 2. 封面圖 3. 原始資料
*/
type InfoWebRawData struct {
	gorm.Model
	InfoStrings []string
	LongTitle   string
	Cover       string
	RawHtml     string
}
