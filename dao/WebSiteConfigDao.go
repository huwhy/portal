package dao

import (
	"github.com/astaxie/beego/orm"
	"portal/model"
	"github.com/astaxie/beego"
)

type WebSiteConfigDao struct {
	ormer orm.Ormer
}

func (dao *WebSiteConfigDao) init() {
	if dao.ormer == nil {
		dao.ormer = orm.NewOrm()
	}
}

func (dao *WebSiteConfigDao) GetSystemConfig() *model.SystemConfig {
	dao.init()
	systemConfig := new(model.SystemConfig)
	systemConfig.CompanyCn = dao.Get("company_cn")
	systemConfig.CompanyEn = dao.Get("company_en")
	systemConfig.CompanyPerson = dao.Get("company_person")
	systemConfig.CompanyPhone = dao.Get("company_phone")
	systemConfig.CompanyMail = dao.Get("company_mail")
	systemConfig.CompanyAddress = dao.Get("company_address")
	return systemConfig
}

func (dao *WebSiteConfigDao) Get(name string) string {
	dao.init()
	var item model.WebSiteConfig
	dao.ormer.Raw("select `name`, content from website_config where `name`=?", name).QueryRow(&item)
	return item.Content
}
func (dao *WebSiteConfigDao) SaveConfig(config model.SystemConfig) {
	dao.init()
	sql := "insert into website_config(name, content) values (?, ?), (?, ?), (?, ?), (?, ?), (?, ?), (?, ?)" +
		" on duplicate key update content=values(content)"

	_, err := dao.ormer.Raw(sql, "company_cn", config.CompanyCn, "company_en", config.CompanyEn,
		"company_person", config.CompanyPerson, "company_phone", config.CompanyPhone, "company_mail", config.CompanyMail, "company_address", config.CompanyAddress).Exec()
	if err != nil {
		beego.Error(err)
	}
}
