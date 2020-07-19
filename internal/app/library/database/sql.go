package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// DB database.
type DB struct {
	write  *conn
	read   []*conn
	idx    int64
	master *DB
}

// Master return *DB instance direct use master conn
// use this *DB instance only when you have some reason need to get result without any delay.
func (db *DB) Master() *DB {
	if db.master == nil {
		panic("ErrNoMaster")
	}
	return db.master
}

func (db *DB) Write() *gorm.DB  {
	return db.write.DB
}

// conn database connection
type conn struct {
	*gorm.DB
	conf    *Config
}

func Open(c *Config) (*DB, error) {
	db := new(DB)
	d, err := connect(c, c.DSN)
	if err != nil {
		return nil, err
	}
	w := &conn{DB: d, conf: c}
	rs := make([]*conn, 0, len(c.ReadDSN))
	for _, rd := range c.ReadDSN {
		d, err := connect(c, rd)
		if err != nil {
			return nil, err
		}
		r := &conn{DB: d, conf: c}
		rs = append(rs, r)
	}
	db.write = w
	db.read = rs
	db.master = &DB{write: db.write}
	return db, nil
}

func connect(c *Config, dataSourceName string) (*gorm.DB, error) {
	d, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	d.DB().SetMaxOpenConns(c.Active)
	d.DB().SetMaxIdleConns(c.Idle)
	d.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	// 不要默认创建数据表添加s后缀
	d.SingularTable(true)
	return d, nil
}
