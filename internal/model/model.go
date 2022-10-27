package model

import (
	miniapp "app-instance/internal/mini-app"
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"gorm.io/gorm"
)

type WhereParam struct {
	Field   string
	Tag     string
	Prepare interface{}
}

type QueryParam struct {
	Fields string
	Offset int
	Limit  int
	Order  string
	Where  []WhereParam
}

func Count(model interface{}, count *int, query QueryParam) bool {
	db := miniapp.App.DB.DbHandler.Model(model)
	db = parseWhereParam(db, query.Where)
	c := (*int64)(unsafe.Pointer(count)) //强转有可能不安全
	db = db.Debug().Count(c)
	if err := db.Error; err != nil {
		miniapp.App.Logger.Warning("database query error: %s", err.Error())
		return false
	}
	return true
}

func parseWhereParam(db *gorm.DB, where []WhereParam) *gorm.DB {
	if len(where) == 0 {
		return db
	}
	var (
		plain   []string
		prepare []interface{}
	)
	for _, w := range where {
		tag := w.Tag
		if tag == "" {
			tag = "="
		}
		var plainFmt string
		switch tag {
		case "IN":
			plainFmt = fmt.Sprintf("%s IN (?)", w.Field)
		default:
			plainFmt = fmt.Sprintf("%s %s ?", w.Field, tag)
		}
		plain = append(plain, plainFmt)
		prepare = append(prepare, w.Prepare)
	}
	// println(strings.Join(plain, " AND "))
	return db.Where(strings.Join(plain, " AND "), prepare...)
}

func Create(model interface{}) bool {
	db := miniapp.App.DB.DbHandler.Debug().Create(model)
	if err := db.Error; err != nil {
		miniapp.App.Logger.Error("database execute error: %s", err.Error())
		return false
	}
	return true
}

func GetOne(model interface{}, query QueryParam) bool {
	db := miniapp.App.DB.DbHandler.Model(model)
	if query.Fields != "" {
		db = db.Select(query.Fields)
	}
	db = parseWhereParam(db, query.Where)
	db = db.First(model)
	if err := db.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		miniapp.App.Logger.Warning("database query error: %s", err.Error())
		return false
	}
	return true
}
