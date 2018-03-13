package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
)

type WebSiteNavDao struct {
	Dao orm.Ormer
}

func (this *WebSiteNavDao) Get(id int) model.WebSiteNav {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var item model.WebSiteNav
	this.Dao.Raw("select `id`, `name` from website_nav where `id`=?", id).QueryRow(&item)
	return item
}
