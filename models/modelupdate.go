package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db2 *gorm.DB

func init() {
	//创建一个数据库的连接
	var err error
	db2, err = gorm.Open("sqlite3", "test222.db")
	if err != nil {
		panic("failed to connect database")
	}

	//迁移the schema
	db2.AutoMigrate(&Pope{})
}

type Pope struct {
	Name       string `json:"name"`       // 登录用户
	Password   string `json:"password"`   // 登录密码
	Permission int    `json:"permission"` // 权限
}

func main() {
	x := []string{"id"}
	data := make(map[string]interface{})
	data["password"] = "77777"
	data["id"] = 2
	pp := &Pope{}
	for k, v := range data {
		fmt.Println(k, v)
		if k != "id" {
			x = append(x, k)
		}
	}
	fmt.Println(x)
	db2.Table("popes").Select("name, password").Where("id = ?", data["id"]).Scan(&pp)
	fmt.Println(*pp)
	//db2.Model(&pp).Where("id = ?", data["id"]).Updates(data)
}
