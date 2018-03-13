package controller

import (
	"portal/model"
	"strconv"
	"encoding/json"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Post() {
	c.init()
	var article model.Article
	json.Unmarshal(c.Ctx.Input.RequestBody, &article)
	id := c.articleDao.Save(article)
	seo := *article.Seo
	seo.Id = id
	c.seoDao.Save(seo)
	c.Data["json"] = model.JsonOk("", nil)
	c.ServeJSON()
}

func (c *ArticleController) List() {
	c.init()
	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 10)
	cid, _ := c.GetInt("cid", 1)
	title := c.GetString("title")
	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = cid
	articleTerm.Page = page
	articleTerm.Size = size
	articleTerm.Title = title
	c.articleDao.Find(articleTerm)
	c.Data["json"] = model.JsonOk("", articleTerm)
	c.ServeJSON()
}

func (c *ArticleController) Article() {
	c.init()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article := c.articleDao.Get(id)
	seo := c.seoDao.Get(id)
	article.Seo = &seo
	c.Data["json"] = model.JsonOk("", article)
	c.ServeJSON()
}

func (c *ArticleController) Get() {
	c.init()
	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone
	c.Data["systemConfig"] = systemConfig

	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 10
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article := c.articleDao.Get(id)
	seo := c.seoDao.Get(id)
	article.Seo = &seo
	c.Data["seo"] = seo

	c.Data["article"] = article
	c.Layout = "common/layout.html"
	c.TplName = "article.html"
}

func (c *ArticleController) ServiceList() {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["seo"] = seo

	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone
	c.Data["systemConfig"] = systemConfig

	c.Data["banners"] = c.bannerDao.GetAll()

	c.Data["title"] = "服务项目"
	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 10
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data
	c.Data["list"] = articleTerm.Data

	c.Layout = "common/layout.html"
	c.TplName = "list.html"
}

func (c *ArticleController) ExampleList() {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["seo"] = seo

	systemConfig := c.configDao.GetSystemConfig()
	c.Data["websiteCn"] = systemConfig.CompanyCn
	c.Data["websiteEn"] = systemConfig.CompanyEn
	c.Data["websitePhone"] = systemConfig.CompanyPhone
	c.Data["systemConfig"] = systemConfig

	c.Data["banners"] = c.bannerDao.GetAll()

	c.Data["title"] = "工程案例"
	articleTerm := new(model.ArticleTerm)
	articleTerm.CatId = 1
	articleTerm.Page = 1
	articleTerm.Size = 10
	c.articleDao.Find(articleTerm)
	c.Data["services"] = articleTerm.Data
	exampleTerm := new(model.ArticleTerm)
	exampleTerm.CatId = 2
	exampleTerm.Page = 1
	exampleTerm.Size = 10
	c.articleDao.Find(exampleTerm)
	c.Data["list"] = exampleTerm.Data

	c.Layout = "common/layout.html"
	c.TplName = "list.html"
}
