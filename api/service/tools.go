package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetListParams(c *gin.Context) (string, []int, []string) {
	var _range []int
	json.Unmarshal([]byte(c.Query("range")), &_range)
	var sort []string
	json.Unmarshal([]byte(c.Query("sort")), &sort)

	return c.Query("filter"), _range, sort
}

func SetContentRange(c *gin.Context, name string, start int, end int, count int) {
	contentRange := fmt.Sprintf("%s %d-%d/%d", name, start, end, count)
	c.Header("Content-Range", contentRange)
}
