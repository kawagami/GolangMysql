package main

import (
	"fmt"
	"mods/sqlGorm"
	"time"
)

type Result struct {
	VName string
	AName string
	LName string
}

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	db := sqlGorm.GetDb()
	//
	var v sqlGorm.Video
	// var a sqlGorm.Actress
	// var s sqlGorm.Storage
	//
	// db.AutoMigrate(&sqlGorm.Video{})
	// db.AutoMigrate(&sqlGorm.Actress{})
	// db.AutoMigrate(&sqlGorm.Storage{})
	// 新增
	// db.Create(&sqlGorm.Video{Name: "test-990"})
	// db.Create(&sqlGorm.Video{Name: "test-001", ActressId: 1})
	// db.Create(&sqlGorm.Video{Name: "test-002", ActressId: 2})
	// db.Create(&sqlGorm.Video{Name: "test-003", ActressId: 3})
	// db.Create(&sqlGorm.Actress{Name: "夢乃あいか"})
	// db.Create(&sqlGorm.Actress{Name: "伊藤舞雪"})
	// db.Create(&sqlGorm.Actress{Name: "さつき芽衣"})
	//
	// 取得第一個
	// if err := db.Where("name = ?", "夢乃あいか").First(&a).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Println(a.ID)
	// fmt.Println(a.Name)
	//
	//
	// 取得第一個
	db.First(&v)
	//
	// var as []sqlGorm.Actress
	var res []Result
	db.Table("videos").Select("videos.name as v_name, actresses.name as a_name, storages.name as l_name").Joins("left join actresses on videos.actress_id = actresses.id").Joins("left join storages on videos.storage_id = storages.id").Find(&res)
	// fmt.Println(res)
	for _, re := range res {
		fmt.Println(re.VName)
		fmt.Println(re.AName)
		fmt.Println(re.LName)
		fmt.Println("")
	}
	//
	// db.First(&a)
	// // db.Model(&v).Association("ActressId").Find(&a)
	// db.Model(&a).Association("Videos").Find(&v)
	// fmt.Println(v)
	//
	// var videos []sqlGorm.Video
	// db.Preload("Actress").Find(&videos)
	// fmt.Println(videos)
	//
	// var actresses []sqlGorm.Actress
	// // 取得所有 actresses 的 videos
	// db.Preload("Videos").Find(&actresses)
	// // fmt.Println(actresses)
	// for _, actress := range actresses {
	// 	fmt.Println(actress.Name)
	// 	for _, video := range actress.Videos {
	// 		fmt.Println(video.Name)
	// 	}
	// 	fmt.Println("")
	// }
	//
	// fmt.Println(v)
	// a.Name = "夢乃あいか"
	// db.Model(&v).Association("Actresses").Append([]sqlGorm.Actress{a})
	// 取得 Actresses 關係的 vs
	// var vs []sqlGorm.Actress
	// db.Model(&v).Association("Actresses").Find(&vs)
	// fmt.Println(vs)
	//
	// var vs []sqlGorm.Actress
	// db.Preload("Actresses").Find(&vs)
	// fmt.Println(vs)
	//
	// var vs []sqlGorm.Video
	// db.Model(&a).Association("Videos").Find(&vs)
	// db.Preload("Actresses")
	// fmt.Println(vs)
	//
	// var vas []sqlGorm.VideoActress
	// db.Find(&vas)
	// for _, va := range vas {
	// 	fmt.Println(va)
	// }
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
