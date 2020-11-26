package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Product struct {
	gorm.Model
	Code string
	Price uint
}
func main() {
	db,err:=Connect()
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	// 自动迁移模式
	db.AutoMigrate(&Product{})
	// 插入
	insert(db)
	// 读取
	//var product Product
	//read(db,&product)
	//fmt.Println(product)
	deleteRow(db,"id=?",15)
	delete(db,&Product{})
}