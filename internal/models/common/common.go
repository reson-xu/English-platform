package common

type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PageResult struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalItems int `json:"totalItems"`
}
