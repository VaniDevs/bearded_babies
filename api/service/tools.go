package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetListParams(c *gin.Context) (string, []int, []string) {
    var _range []int
	json.Unmarshal([]byte(c.Query("range")), &_range)
	var sort []string
	json.Unmarshal([]byte(c.Query("sort")), &sort)
	var _filter map[string]interface{}
    json.Unmarshal([]byte(c.Query("filter")), &_filter)
    var filterString = ""
    if len(_filter) > 0 {
        var where = make([]string, 0, len(_filter))
        for key, value := range _filter {
            v, ok := value.(float64)
            if ok {
                where = append(where, fmt.Sprintf(" %s=%v ", key, v))
            } else {
                v1, ok1 := value.([]interface{})
                if ok1 {
                    fmt.Println(ok1)
                    replacer := strings.NewReplacer("[", "", "]", "")
                    where = append(where, replacer.Replace(fmt.Sprintf(" %s in (%v) ", key, v1)))
                } else {
                    where = append(where, fmt.Sprintf(" %s='%v' ", key, value))
                }
            }
        }
        filterString = strings.Join(where[:], " and ")
    }

	return filterString, _range, sort
}

func SetContentRange(c *gin.Context, name string, start int, end int, count int) {
	contentRange := fmt.Sprintf("%s %d-%d/%d", name, start, end, count)
	c.Header("Content-Range", contentRange)
}
