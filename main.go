package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"portal/model"
	"portal/controller"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("data.source"), 30)
	//orm.RegisterModel(new(model.Article))
	//orm.RunSyncdb("default", false, false)
}

var LoginFilter = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("sc_id").(int)
	if !ok && ctx.Request.RequestURI != "/api/login" {
		result := model.Json302("/login")
		bytes, _ := json.Marshal(result)
		ctx.ResponseWriter.Write(bytes)
	}
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Router("/", &controller.IndexController{})
	aboutController := &controller.AboutController{}
	beego.Router("/about.html", aboutController)
	beego.Router("/contact.html", aboutController, "Get:Contact")
	articleController := &controller.ArticleController{}
	beego.Router("/services.html", articleController, "Get:ServiceList")
	beego.Router("/examples.html", articleController, "Get:ExampleList")
	beego.Router("/article/:id([0-9]+).html", articleController)
	beego.Router("/file", &controller.FileController{})

	//api
	beego.Router("/api/article", articleController, "Get:List")
	beego.Router("/api/article", articleController)
	beego.Router("/api/article/:id([0-9]+)", articleController, "Get:Article")
	beego.Router("/api/upload", &controller.FileController{})
	bannerController := &controller.BannerController{}
	beego.Router("/api/banner", bannerController)
	beego.Router("/api/banner/:id([0-9]+)", bannerController)
	configController := &controller.SystemConfigController{}
	beego.Router("/api/system", configController)
	beego.Router("/api/seo", configController, "Get:GetSeo")
	beego.Router("/api/seo", configController, "Post:SaveSeo")
	beego.SetStaticPath("/static", "static")
	beego.SetLogger("file", `{"filename":"logs/app.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	//beego.BConfig.WebConfig.TemplateLeft = "<%"
	//beego.BConfig.WebConfig.TemplateRight = "%>"
	beego.Run()
}
