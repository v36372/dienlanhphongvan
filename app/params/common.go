package params

import "strconv"

const (
	paramUrlSlug  = "slug"
	paramUrlLimit = "limit"
	paramUrlPage  = "page"

	maxLimit     = 100
	defaultLimit = 20
	defaultPage  = 1
)

func parseUrlParamToInt(value string, defaultVal int) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}

	return res
}
