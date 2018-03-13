package controller

import (
	"portal/model"
	"encoding/json"
)

type SystemConfigController struct {
	BaseController
}

func (c *SystemConfigController) Get() {
	c.init()
	c.Data["json"] = model.JsonOk("", c.configDao.GetSystemConfig())
	c.ServeJSON()
}

func (c *SystemConfigController) GetSeo () {
	c.init()
	seo := c.seoDao.Get(0)
	c.Data["json"] = model.JsonOk("", seo)
	c.ServeJSON()
}

func (c *SystemConfigController) SaveSeo () {
	c.init()
	var seo model.Seo
	json.Unmarshal(c.Ctx.Input.RequestBody, &seo)
	c.seoDao.Save(seo)
	c.Data["json"] = model.JsonOk("", nil)
	c.ServeJSON()
}

func (c *SystemConfigController) Post() {
	c.init()
	var systemConfig model.SystemConfig
	json.Unmarshal(c.Ctx.Input.RequestBody, &systemConfig)
	c.configDao.SaveConfig(systemConfig)
	c.Data["json"] = model.JsonOk("", nil)
	c.ServeJSON()
}

