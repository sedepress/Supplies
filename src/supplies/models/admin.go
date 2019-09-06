package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Admin))
}

type Admin struct {
	Id int `json:"id" orm:"pk;unique;auto_increment;column(id)"`
	Name string `json:"name" orm:"unique;size(32);column(name)"`
	Password string `json:"password" orm:"size(128);column(password)"`
	NickName string `json:"nick_name" orm:"index;size(32);column(nick_name)"`
	Salt string `json:"salt" orm:"size(128);column(salt)"`
	Token string `json:"token" orm:"null;column(token)"`
	LastLogin time.Time `json:"last_login" orm:"null;column(last_login)"`
	Status int `json:"status" orm:"size(1);default(1);column(status)"`
	CreateAt time.Time `json:"create_at" orm:"auto_now_add;type(datetime);column(create_at)"`
	UpdateAt time.Time `json:"update_at" orm:"auto_now;type(datetime);column(update_at)"`
}

func Admins() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Admin))
}