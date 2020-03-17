package mysqlTools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"sync"
	"time"
)

type MysqlPool struct {

}

var instance *MysqlPool
var once sync.Once

var db *gorm.DB
var errorDb error

func GetInstance() *MysqlPool {
	once.Do(func() {
		instance = &MysqlPool{}
	})
	return instance
}

func (m *MysqlPool) InitDataPool() (status bool) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")
	str := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, dbName)
	fmt.Println(str)
	db, errorDb = gorm.Open("mysql", str)
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	fmt.Println(errorDb)
	if errorDb != nil {
		log.Fatal(errorDb)
		return false
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true
}

func (m *MysqlPool) GetMysqlDB() (_db *gorm.DB) {
	return db
}

