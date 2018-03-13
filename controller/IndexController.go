package controller

import (
	"portal/model"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["seo"] = seo

	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone

	c.Data["banners"] = c.bannerDao.GetAll()

	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 100
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data

	summary := c.articleDao.Get(1)
	c.Data["companySummary"] = summary.Content

	exampleTerm := new(model.ArticleTerm)
	exampleTerm.CatId = 2
	exampleTerm.Page = 1
	exampleTerm.Size = 10
	c.articleDao.Find(exampleTerm)
	c.Data["examples"] = exampleTerm.Data

	knowTerm := new(model.ArticleTerm)
	knowTerm.CatId = 3
	knowTerm.Page = 1
	knowTerm.Size = 10
	c.articleDao.Find(knowTerm)
	c.Data["knows"] = knowTerm.Data

	c.Layout = "common/layout.html"
	c.TplName = "index.html"
}
