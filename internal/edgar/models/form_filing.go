package models

type FormFiling struct {
	CIK         string
	CompanyName string
	Type        FormType
	DateFiled   string
	FileName    string
}
