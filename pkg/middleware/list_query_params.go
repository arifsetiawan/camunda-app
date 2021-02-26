package middleware

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/arifsetiawan/camunda-app/pkg/queryparams"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`\[(.*?)\]`)
}

// ListQueryParams middleware get jsonapi defined paging, sort and filter query
func ListQueryParams(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := c.QueryParams()

		pageNumberExist := false
		pageOffsetExist := false

		listQueryParams := new(queryparams.ListQueryParams)
		listQueryParams.Filter = make(map[string]string)

		for k := range queryParams {
			//fmt.Printf("key[%s] value[%s]\n", k, v)
			if k == "page[number]" {
				listQueryParams.PageNumber, _ = strconv.Atoi(c.QueryParam("page[number]"))
				pageNumberExist = true
			}
			if k == "page[size]" {
				listQueryParams.PageSize, _ = strconv.Atoi(c.QueryParam("page[size]"))
			}
			if k == "page[offset]" {
				listQueryParams.PageOffset, _ = strconv.Atoi(c.QueryParam("page[offset]"))
				pageOffsetExist = true
			}
			if k == "page[limit]" {
				listQueryParams.PageLimit, _ = strconv.Atoi(c.QueryParam("page[limit]"))
			}
			if strings.Contains(k, "filter") {
				match := re.FindStringSubmatch(k)
				if len(match) == 2 {
					listQueryParams.Filter[match[1]] = c.QueryParam(k)
				}
			}
			if strings.Contains(k, "sort") {
				listQueryParams.Sort = strings.Split(c.QueryParam(k), ",")
			}
		}

		// if we have page-size we convert it into offset-limit
		if pageNumberExist && !pageOffsetExist {
			// boundary check
			if listQueryParams.PageNumber < 1 {
				listQueryParams.PageNumber = 1
			}

			listQueryParams.PageOffset = (listQueryParams.PageNumber - 1) * listQueryParams.PageSize
			listQueryParams.PageLimit = listQueryParams.PageSize
		}

		// if we have offset-limit we convert it into page-size
		if !pageNumberExist && pageOffsetExist {
			listQueryParams.PageNumber = listQueryParams.PageOffset/listQueryParams.PageLimit + 1

			// boundary check
			if listQueryParams.PageNumber < 1 {
				listQueryParams.PageNumber = 1
			}

			listQueryParams.PageSize = listQueryParams.PageLimit
		}

		// default
		if listQueryParams.PageSize == 0 {
			listQueryParams.PageSize = 10
		}

		if listQueryParams.PageLimit == 0 {
			listQueryParams.PageLimit = 10
		}

		c.Set("listQueryParams", listQueryParams)
		return next(c)
	}
}
