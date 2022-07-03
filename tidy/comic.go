package tidy

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 取得 & 整理作者名 slice
func ComicAuthors() (res []Tidy) {
	path := `C:\comic\H`
	// path := `D:\comic\H\`
	directories := GetFilesNameOfDir(path)
	//
	var tempAuthors []Tidy
	pattern := `\[(.*)\]`
	re := regexp.MustCompile(pattern)
	for _, directory := range directories {
		if find := re.FindStringSubmatch(directory.Title); len(find) > 0 {
			// fmt.Println("len = ", len(find))
			// fmt.Printf("find[0] = %v\n", find[0])
			// fmt.Printf("find[1] = %v\n", find[1])
			var data Tidy
			data.Title = find[1]
			data.Path = directory.Path
			tempAuthors = append(tempAuthors, data)
		}
	}
	//
	pattern2 := `(.*)\s?\((.*)\)`
	re2 := regexp.MustCompile(pattern2)
	for _, Author := range tempAuthors {
		if find2 := re2.FindStringSubmatch(Author.Title); len(find2) > 0 {
			// fmt.Println("len =", len(find2))
			// fmt.Printf("find[0] = %v\n", find2[0])
			// fmt.Printf("find[1] = %v\n", find2[1])
			// fmt.Printf("find[2] = %v\n", find2[2])
			// fmt.Println("path =", Author.Path)
			var name1 Tidy
			name1.Title = find2[1]
			name1.Path = Author.Path
			var name2 Tidy
			name2.Title = find2[2]
			name2.Path = Author.Path
			res = append(res, name1)
			res = append(res, name2)
		} else {
			var data Tidy
			data.Title = Author.Title
			data.Path = Author.Path
			res = append(res, data)
		}
	}
	return
}

// 取得 temp 下所有檔名 & 路徑
func TempComics() (res []Tidy) {
	path := `D:\temp\`
	res = GetFilesNameOfDir(path)
	// for _, comic := range comics {
	// 	res = append(res, comic)
	// 	// var data Tidy
	// 	// fmt.Println(comic.Title)
	// 	// fmt.Println(comic.Path)
	// 	// fmt.Println("")
	// }
	//
	return
}

/*
綜合使用 ComicAuthors() TempComics()
完成整理 temp 資料夾的目標
移動 temp 資料夾的壓縮檔到同作者名的資料夾下
*/
func MoveFileToDest() {
	count := 0
	authors := ComicAuthors()
	comics := TempComics()
	for _, author := range authors {
		for _, comic := range comics {
			if strings.Contains(comic.Path, author.Title) {
				needCheckFileExist := author.Path + `\` + comic.Title
				comicPathCheck, statErr := os.Stat(comic.Path)
				if statErr != nil {
					fmt.Println("statErr")
					fmt.Println("")
					continue
				}
				comicPathIsFile := !comicPathCheck.IsDir()
				if _, err := os.Stat(needCheckFileExist); os.IsNotExist(err) && comicPathIsFile {
					count++
					fmt.Println(count)
					fmt.Println(comic.Path)
					fmt.Println("moving to")
					fmt.Println(needCheckFileExist)
					fmt.Println(comicPathIsFile)
					fmt.Println("")
					os.Rename(comic.Path, needCheckFileExist)
				} else {
					fmt.Println("檔案已經存在", needCheckFileExist)
					fmt.Println("")
				}
			}
		}
	}
	fmt.Println(count)
}
