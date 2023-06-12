package untils

import (
	"regexp"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// influxdb 查询结果处理
func InfluxdbQueryResult(result *api.QueryTableResult) ([]map[string]interface{}, error){
	results := []map[string]interface{}{}
	for result.Next() {
		record := result.Record()
		item := map[string]interface{}{}
		for key, value := range record.Values() {
			item[key] = value
		}
		results = append(results, item)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
