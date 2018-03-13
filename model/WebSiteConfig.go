package model

type WebSiteConfig struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type SystemConfig struct {
	CompanyCn string `json:"companyCn"`
	CompanyEn string `json:"companyEn"`
	CompanyPerson string `json:"companyPerson"`
	CompanyPhone string `json:"companyPhone"`
	CompanyMail string `json:"companyMail"`
	CompanyAddress string `json:"companyAddress"`
}
