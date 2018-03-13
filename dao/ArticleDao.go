package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
	"fmt"
	"time"
	"github.com/astaxie/beego"
)

type ArticleDao struct {
	ormer orm.Ormer
}

func (dao *ArticleDao) init() {
	if dao.ormer == nil {
		dao.ormer = orm.NewOrm()
	}
}

func (dao *ArticleDao) Get(id int) model.Article {
	dao.init()
	var item model.Article
	dao.ormer.Raw("select `id`, cat_id, title, main_img, content, view_num, created from article where `id`=?", id).QueryRow(&item)
	return item
}

func (dao *ArticleDao) Find(term *model.ArticleTerm) {
	dao.init()
	var args []interface{}
	var articles []model.Article
	sql := "select `id`, cat_id, title, main_img, view_num, created from article where 1=1"
	if term.CatId != 0 {
		sql += " and cat_id=?"
		args = append(args, term.CatId)
	}
	if len(term.Title) > 0 {
		sql += " and title like ?"
		args = append(args, "%"+term.Title+"%")
	}

	id, err := Paging(sql, "", args, &term.Term, dao.ormer).QueryRows(&articles)
	if err != nil {
		panic(err)
		fmt.Println(id)
	}
	term.Data = articles
}

func (dao *ArticleDao) Save(article model.Article) int {
	dao.init()
	article.ViewNum = 0
	article.Created = time.Now()
	inserted := article.Id == 0
	r, err := dao.ormer.Raw(`insert into article(id, cat_id, title, main_img, content, view_num, created)
		values (?, ?, ?, ?, ?, ?, ?) on duplicate key update cat_id=values(cat_id), title=values(title),
		main_img=values(main_img), content=values(content)
		`, article.Id, article.CatId, article.Title, article.MainImg, article.Content,
		article.ViewNum, article.Created).Exec()
	if err != nil {
		beego.Error(err)
	} else if (inserted){
		id, _ := r.LastInsertId()
		article.Id = int(id)
	}
	return article.Id
}
