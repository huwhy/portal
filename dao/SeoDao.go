package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
	"github.com/astaxie/beego"
)

type SeoDao struct {
	ormer orm.Ormer
}

func (dao *SeoDao) init() {
	if dao.ormer == nil {
		dao.ormer = orm.NewOrm()
	}
}

func (dao *SeoDao) Get(id int) model.Seo {
	dao.init()
	var seo model.Seo
	dao.ormer.Raw("select `id`, `title`, `desc`, `words` from seo where `id`=?", id).QueryRow(&seo)
	return seo
}

func (dao *SeoDao) Save(seo model.Seo) {
	dao.init()
	sql := "insert into seo(id, title, `desc`, words) values (?, ?, ?, ?) " +
		"on duplicate key update title=values(title), `desc`=values(`desc`), words=values(words)"
	_, err := dao.ormer.Raw(sql, seo.Id, seo.Title, seo.Desc, seo.Words).Exec()
	if err != nil {
		beego.Error(err)
	}
}

func (dao *SeoDao) Del(id int) {
	dao.init()
	_, err := dao.ormer.Raw("delete from seo where id=?", id).Exec();
	if err != nil {
		panic(err)
	}
}
