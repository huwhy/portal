package controller

import (
	"encoding/json"
	"portal/model"
	"strconv"
)

type BannerController struct {
	BaseController
}

func (c *BannerController) Post() {
	c.init()
	var banner model.WebSiteBanner
	json.Unmarshal(c.Ctx.Input.RequestBody, &banner)
	c.bannerDao.Save(banner)
	c.Data["json"] = model.JsonOk("", nil)
	c.ServeJSON()
}

func (c *BannerController) Get() {
	c.init()
	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 10)
	term := new(model.Term)
	term.Page = page
	term.Size = size
	c.bannerDao.Find(term)
	c.Data["json"] = model.JsonOk("", term)
	c.ServeJSON()
}

func (c *BannerController) Delete() {
	c.init()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.bannerDao.Del(id)
	c.Data["json"] = model.JsonOk("", nil)
	c.ServeJSON()
}

