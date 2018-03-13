package model

import "time"

type Article struct {
	Id int `json:"id"`
	CatId int `json:"catId"`
	Title string `json:"title"`
	MainImg string `json:"mainImg"`
	Content string `json:"content"`
	ViewNum int `json:"viewNum"`
	Created time.Time `json:"created"`
	Seo *Seo `json:"seo"`
}

type ArticleTerm struct {
	Term
	CatId int `json:"catId"`
	Title string `json:"title"`
}