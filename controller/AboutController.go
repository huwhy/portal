package controller

import "portal/model"

type AboutController struct {
	BaseController
}

func (c *AboutController) Get() {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["seo"] = seo
	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone
	c.Data["systemConfig"] = systemConfig

	c.Data["banners"] = c.bannerDao.GetAll()

	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 100
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data

	summary := c.articleDao.Get(1)
	c.Data["companySummary"] = summary.Content

	c.Layout = "common/layout.html"
	c.TplName = "about.html"
}

func (c *AboutController) Contact() {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["seo"] = seo

	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone
	c.Data["systemConfig"] = systemConfig

	c.Data["banners"] = c.bannerDao.GetAll()

	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 100
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data

	businessScope := c.articleDao.Get(0)
	c.Data["businessScope"] = businessScope.Content

	c.Layout = "common/layout.html"
	c.TplName = "contact.html"
}
