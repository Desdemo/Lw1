package main

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// 银行账户
type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

type Datachange interface {
	Createer()
	Updateer()
	Queryer()
	Delecter()
}

// ORM 引擎
var engine *xorm.Engine

func init() {
	// 创建 ORM 引擎与数据库
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./te.db")
	engine.ShowSQL(true)
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = engine.Sync2(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func (a Account) Creater() {
	_, err := engine.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 创建新的账户
func newAccount(name string, balance float64) error {
	// 对未存在记录进行插入
	_, err := engine.Insert(&Account{Name: name, Balance: balance})
	return err
}

func getAccount(id int64) (*Account, error) {
	a := &Account{}
	// 直接操作 ID 的简便方法
	has, err := engine.Id(id).Get(a)
	// 判断操作是否发生错误或对象是否存在
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account does not exist")
	}
	return a, nil
}

func main() {

	fmt.Println((getAccount(3)))

}
