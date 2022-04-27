package GetInfo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
遍歷所有層的資料夾和檔案
*/
func GetAllFile(files *[]string) filepath.WalkFunc {
	/* 範例
	var files []string
	root := "D:\\temp"
	err := filepath.Walk(root, GetInfo.GetAllFile(&files))
	if err != nil {
		panic(err)
	}
	*/
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("visit 的 path = %v\n", path)
		*files = append(*files, path)
		return nil
	}
}

/*
只遍歷下一層的所有資料夾
*/
func GetDir(pathname string, pathSlice *[]string) error {
	var tempSlice []string
	//
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			// fmt.Printf("是資料夾 %s\n", pathname+"\\"+fi.Name())
			// fmt.Printf("資料夾名稱 %s\n", fi.Name())
			// fmt.Println()
			//
			tempSlice = append(tempSlice, pathname+"\\"+fi.Name())
			fmt.Println("get", fi.Name())
			// GetAllFile(pathname + fi.Name() + "\\")
		} else {
			// fmt.Println(fi.Name())
		}
	}
	*pathSlice = tempSlice
	return err
}
