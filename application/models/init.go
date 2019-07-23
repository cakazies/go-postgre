package models

import (
	"fmt"
	"strconv"
	"strings"
)

func LimitOffset(limit string, offset string) string {
	limits, _ := strconv.Atoi(limit)
	offsets, _ := strconv.Atoi(offset)
	if limit == "" && offsets > 0 {
		return fmt.Sprintf(" OFFSET %s", offset)
	}
	if offset == "" && limits > 0 {
		return fmt.Sprintf(" LIMIT %s", limit)
	}
	if limits > 0 && offsets > 0 {
		limit := strconv.Itoa(limits)
		offset := strconv.Itoa(offsets)
		return fmt.Sprintf(" LIMIT %s OFFSET %s", limit, offset)
	}
	return ""
}

func ShortBy(query string) string {
	result := ""
	if query == "" {
		return " ORDER BY rm_id desc "
	}
	arrQuery := strings.Split(query, ",")
	for k, v := range arrQuery {
		value := strings.Split(v, ".")
		if k == 0 {
			result = fmt.Sprintf(" ORDER BY %s %s ", value[0], value[1])
		}
		if k > 0 {
			result += fmt.Sprintf(" , %s %s ", value[0], value[1])
		}
	}
	return result
}
