package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
)

type ArticleCategoryDao struct {
	Dao orm.Ormer
}

func (this *ArticleCategoryDao) Get(id int) model.ArticleCategory {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var item model.ArticleCategory
	this.Dao.Raw("select `id`, `name` from article_category where `id`=?", id).QueryRow(&item)
	return item
}
