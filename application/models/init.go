package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	InitViper()
	Connect()
}

// InitViper from file toml
func InitViper() {
	viper.SetConfigFile("toml")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("viper not use")
	}
}

// LimitOffset function for get limit and offset in request API
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

// ShortBy function for shorting data
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
