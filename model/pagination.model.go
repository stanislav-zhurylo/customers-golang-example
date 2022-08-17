package model

type PaginationSettings struct {
	Page          int    `schema:"page"`
	PageSize      int    `schema:"pageSize"`
	Sort          string `schema:"sort"`
	Direction     string `schema:"direction"`
	TotalElements int    `schema:"totalElements"`
	TotalPages    int    `schema:"totalPages"`
}
