package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


type MysqlCon struct{
	Name            string
	Addr            string
	port            int
	Db              string
	Username        string
	Password        string
	MaxIdealConn    int
	MaxOpenConn     int
	ConnMaxLifetime int
}
var mysqlCon =MysqlCon{
	"FilNest",
	"61.174.254.164",
	3306,
	"filnest",
	"filnest",
	"HTB3kxTE32czEe7a",
	10,
	256,
	600,

}
func Connect() (*gorm.DB,error){
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，
	//我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlCon.Username, mysqlCon.Password, mysqlCon.Addr, mysqlCon.port, mysqlCon.Db, "10s")
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open("mysql", dsn)

	return db,err
}