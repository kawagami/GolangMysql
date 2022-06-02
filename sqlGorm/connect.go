package sqlGorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type VideoActress struct {
	gorm.Model
	Title   string
	Actress string
}

func GetDb() (db *gorm.DB) {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	// dbname := "gorm"
	dbname := "arrange"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db

	// // 迁移 schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	// var product Product
	// db.Migrator().DropTable(&product)
	// var product VideoActress
	// var vas []VideoActress
	// db.First(&product) // 根据整型主键查找
	// // db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // fmt.Println(product.Title)
	// // fmt.Println(product.Actress)

	// db.Find(&vas)

	// for _, va := range vas {
	// 	fmt.Println(va.Title)
	// 	fmt.Println(va.Actress)
	// 	fmt.Println("")
	// }

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
