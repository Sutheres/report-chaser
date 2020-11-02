package models

type Item struct {
	LastModified string `json:"last-modified"`
	Name string `json:"name"`
	Type string `json:"type"`
	Href string `json:"href"`
	Size string `json:"size"`
}

type Directory struct {
	Items []Item `json:"item"`
	Name string `json:"name"`
	ParentDirectory string `json:"parent-directory"`
}

type DailyIndex struct {
	Directory `json:"directory"`
}
