package godb

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Unix            string `yaml:"unix"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Charset         string `yaml:"charset"`
	User            string `yaml:"user"`
	Pass            string `yaml:"password"`
	DbName          string `yaml:"dbname"`
	TablePrefix     string `yaml:"tableprefix"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifeTime int    `yaml:"conn_max_life_time"`
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
