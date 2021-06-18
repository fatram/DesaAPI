package main

type Village struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
	Regency  string `json:"regency"`
	Province string `json:"province"`
}

type MetaInfo struct {
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
	CurrentPage  int `json:"current_page"`
	TotalPage    int `json:"total_page"`
}

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Village `json:"data"`
	Meta    MetaInfo  `json:"meta"`
}
