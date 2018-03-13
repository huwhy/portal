package controller

import (
	"github.com/astaxie/beego"
	"portal/dao"
)

type BaseController struct {
	beego.Controller
	configDao  *dao.WebSiteConfigDao
	bannerDao  *dao.WebSiteBannerDao
	articleDao *dao.ArticleDao
	seoDao *dao.SeoDao
}

func (c *BaseController) init() {
	if c.configDao == nil {
		c.configDao = new(dao.WebSiteConfigDao)
		c.bannerDao = new(dao.WebSiteBannerDao)
		c.articleDao = new(dao.ArticleDao)
		c.seoDao = new(dao.SeoDao)
	}
}


