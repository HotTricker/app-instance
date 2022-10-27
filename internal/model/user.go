package model

import (
	"errors"
)

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"type:string;not null;default:''"`
	Password string `gorm:"type:string;not null"`
}

func (u *User) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(u, &count, query)
	return count, ok
}

func (u *User) TableName() string {
	return "stu_user"
}

func (u *User) Create() bool {
	return Create(u)
}

func (u *User) UserCheckExists() (bool, error) {
	var where []WhereParam
	if u.Username != "" {
		where = append(where, WhereParam{
			Field:   "username",
			Prepare: u.Username,
		})
	}
	if u.ID != 0 {
		where = append(where, WhereParam{
			Field:   "id",
			Tag:     "!=",
			Prepare: u.ID,
		})
	}

	count, ok := u.Count(QueryParam{
		Where: where,
	})

	if !ok {
		return false, errors.New("check user failed")
	}
	return count > 0, nil
}

func (u *User) CreateOrUpdate() error {
	if u.ID > 0 {
		//TODO: need update
		ok := u.Create()
		if !ok {
			return errors.New("user update failed")
		}
	} else {
		ok := u.Create()
		if !ok {
			return errors.New("user update failed")
		}
	}
	return nil
}

func (u *User) GetOne(query QueryParam) bool {
	return GetOne(u, query)
}
