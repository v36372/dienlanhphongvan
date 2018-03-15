package view

import "github.com/gin-gonic/gin"

type Response struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"totalPage"`
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(200, Response{Data: data})
}

func ResponseOKWithPagination(c *gin.Context, data interface{}, pagination *Pagination) {
	c.JSON(200, Response{
		Data:       data,
		Pagination: pagination,
	})
}

func NewPagination(total, limit, page int) Pagination {
	return Pagination{
		Total:     total,
		Page:      page,
		Limit:     limit,
		TotalPage: (total / limit) + 1,
	}
}
