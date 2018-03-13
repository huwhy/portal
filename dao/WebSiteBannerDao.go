package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
	"fmt"
	"github.com/astaxie/beego"
)

type WebSiteBannerDao struct {
	ormer orm.Ormer
}

func (dao *WebSiteBannerDao) init() {
	if dao.ormer == nil {
		dao.ormer = orm.NewOrm()
	}
}

func (dao *WebSiteBannerDao) Get(id int) model.WebSiteBanner {
	dao.init()
	var banner model.WebSiteBanner
	dao.ormer.Raw("select `id`, `img`, `url` from website_banner where `id`=?", id).QueryRow(&banner)
	return banner
}

func (dao *WebSiteBannerDao) GetAll() []model.WebSiteBanner {
	dao.init()
	var banners []model.WebSiteBanner
	dao.ormer.Raw("select `id`, `img`, `url` from website_banner").QueryRows(&banners)
	return banners
}

func (dao *WebSiteBannerDao) Find(term *model.Term) {
	dao.init()
	var args []interface{}
	var banners []model.WebSiteBanner
	sql := "select `id`, `img`, `url` from website_banner where 1=1"
	orderSql := " order by id desc"
	id, err := Paging(sql, orderSql, args, term, dao.ormer).QueryRows(&banners)
	if err != nil {
		panic(err)
		fmt.Println(id)
	}
	term.Data = banners
}

func (dao *WebSiteBannerDao) Save(banner model.WebSiteBanner) int {
	dao.init()
	r, err := dao.ormer.Raw(`insert into website_banner(id, img, url)
		values (?, ?, ?)`, banner.Id, banner.Img, banner.Url).Exec()
	if err != nil {
		beego.Error(err)
	} else {
		id, _ := r.LastInsertId()
		banner.Id = int(id)
	}
	return banner.Id
}

func (dao *WebSiteBannerDao) Del(id int) {
	dao.init()
	_, err := dao.ormer.Raw("delete from website_banner where id=?", id).Exec();
	if err != nil {
		panic(err)
	}
}

