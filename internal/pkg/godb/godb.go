package godb

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Unix            string
	Host            string
	Port            int
	Charset         string
	User            string
	Pass            string
	DbName          string
	TablePrefix     string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime int
}

type DB struct {
	DbHandler *gorm.DB
	cfg       *DbConfig
}

func NewDatabase(cfg *DbConfig) *DB {
	return &DB{
		cfg: cfg,
	}
}

func (db *DB) Open() error {
	if c, err := gorm.Open(mysql.Open(db.parseConnConfig())); err == nil {
		if sqlDb, err := c.DB(); err == nil {
			sqlDb.SetMaxIdleConns(db.cfg.MaxIdleConns)
			sqlDb.SetMaxOpenConns(db.cfg.MaxOpenConns)
			sqlDb.SetConnMaxLifetime(time.Second * time.Duration(db.cfg.ConnMaxLifeTime))
			db.DbHandler = c
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (db *DB) parseConnConfig() string {
	var connHost string
	if db.cfg.Unix != "" {
		connHost = fmt.Sprintf("unix(%s)", db.cfg.Unix)
	} else {
		connHost = fmt.Sprintf("tcp(%s:%d)", db.cfg.Host, db.cfg.Port)
	}
	s := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local", db.cfg.User, db.cfg.Pass, connHost, db.cfg.DbName, db.cfg.Charset)
	return s
}
