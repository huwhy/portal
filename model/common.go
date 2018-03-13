package model

type Term struct {
	Size     int         `json:"size"`
	Page     int         `json:"page"`
	Total    int         `json:"total"`
	HasStart bool        `json:"hasStart"`
	HasTotal bool        `json:"hasTotal"`
	Data     interface{} `json:"data"`
}

func (this *Term) Start() int {
	return (this.Page - 1) * this.Size
}

func (this *Term) TotalPage() int {
	return this.Total / this.Size
}

type Json struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Url     string      `json:"url"`
}

func NewJson() Json {
	json := Json{}
	return json
}

func Json302(url string) Json {
	json := Json{Code: 302, Url: url}
	return json
}

func JsonOk(message string, data interface{}) Json {
	return Json{200, message, data, ""}
}

func JsonErr(message string) Json {
	return Json{500, message, nil, ""}
}