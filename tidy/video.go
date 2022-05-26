package tidy

import (
	"fmt"
	"mods/GetInfo"
	"os"
	"path/filepath"
)

// 將 tempVideoPath 的影片移動到 destPath 中有該演員名稱的資料夾中
func TidyVideoByActressName(tempVideoPath, destPath string, move bool) {
	// map[番號]演員名 的 map
	res := GetInfo.GetSqlDataMap()
	// 取得 path 路徑中符合檔案名稱格式的檔名和完整路徑
	// 要整理的資料夾
	// tempVideoPath := `E:\video\H\unArranged\`
	// destPath := `E:\video\H\true\`
	//
	result := GetInfo.GetFileDataFromDir(tempVideoPath)
	count := 0
	for _, vc := range result {
		// res[vc.Title] != nil 的話就代表資料庫有該片的資料
		if res[vc.Title] != "" {
			// 包含副檔名的檔案名稱
			tempFileName := filepath.Base(vc.Path)
			// 移動後的檔案名稱
			finalName := destPath + res[vc.Title] + `\` + tempFileName
			// 檢查 finalName 是否已經存在
			if _, err := os.Stat(finalName); err == nil {
				fmt.Println(finalName)
				fmt.Printf("File exists\n")
				fmt.Println("")
			} else {
				count++
				fmt.Println(vc.Path)
				fmt.Println("要變成")
				fmt.Println(finalName)
				// 目前有目標資料夾不存在時不會自動建立 & 檔名大小寫有差也不會移動的問題
				if move {
					os.Rename(vc.Path, finalName)
				}
				// fmt.Printf("File does not exist\n")
				fmt.Println("")
			}
		}
	}
	fmt.Println(count)
}
