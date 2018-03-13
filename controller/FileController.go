package controller

import (
	"log"
	"portal/model"
	"github.com/astaxie/beego"
	"os"
	"time"
)

type FileController struct {
	BaseController
}

func (c *FileController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	todayStr := time.Now().Format("20060102")
	path := beego.AppConfig.String("upload.dir") + todayStr
	_, err = os.Stat(path)
	if err == nil || os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	err = c.SaveToFile("file", path + "/" + h.Filename)
	if err != nil {
		log.Fatal("savefile err ", err)
	}
	host := beego.AppConfig.String("upload.host")
	c.Data["json"] = model.JsonOk("", host + todayStr + "/" + h.Filename)
	c.ServeJSON()
}

func (c *FileController) Get() {
	filename := c.GetString("file")
	file := beego.AppConfig.String("upload.dir") + filename
	c.Ctx.Output.Download(file, filename)
}