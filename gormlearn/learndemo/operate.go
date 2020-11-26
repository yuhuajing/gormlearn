package main

import "github.com/jinzhu/gorm"

//插入
func insert(db *gorm.DB){
	db.Create(&Product{Code: "L1212", Price: 1000})
}
//查
func read(db *gorm.DB, product *Product){
	db.First(product, 1) // 查询id为1的product
	db.First(product, "code = ?", "L1212")
	db.Find(product)
}
// 更新 - 更新product的price为2000
func modify(db *gorm.DB, product *Product){
	db.Model(product).Update("Price", 2000)
}
//删除指定条件
func deleteRow(db *gorm.DB,str string ,id uint){
	db.Unscoped().Where(str,id).Delete(&Product{})
}

// 删除 - 删除表
func delete(db *gorm.DB, product *Product) {
	db.DropTable(Product{})
}