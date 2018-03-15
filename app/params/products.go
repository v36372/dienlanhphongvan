package params

import "github.com/gin-gonic/gin"

func NewGetProductSlugParam(c *gin.Context) (producSlug string) {
	return c.Param(paramUrlSlug)
}

func NewGetProductsParams(c *gin.Context) (limit int, offset int, page int) {
	limit = parseUrlParamToInt(c.Param(paramUrlLimit), defaultLimit)
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	page = parseUrlParamToInt(c.Param(paramUrlPage), defaultPage)
	if page <= 0 {
		page = defaultPage
	}
	offset = (page - 1) * limit

	return
}
