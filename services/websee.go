package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
)

func serializeField(field interface{}) (string, error) {
	if field == nil {
		return "", nil
	}
	jsonBytes, err := json.Marshal(field)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func SetMonitor(data *models.ReportData) *write.Point {
	dataJson := &models.ReportDataJson{
		Name:           data.Name,
		Type:           data.Type,
		PageUrl:        data.PageUrl,
		Rating:         data.Rating,
		Time:           time.Now().UnixNano(),
		UUID:           data.UUID,
		Apikey:         data.Apikey,
		Status:         data.Status,
		SdkVersion:     data.SdkVersion,
		Events:         data.Events,
		UserId:         data.UserId,
		Line:           data.Line,
		Column:         data.Column,
		Message:        data.Message,
		RecordScreenId: data.RecordScreenId,
		FileName:       data.FileName,
		Url:            data.Url,
		ElapsedTime:    data.ElapsedTime,
		Value:          data.Value,
	}

	if data.Type == "xhr" || data.Type == "fetch" {
		dataJson.Type = "apiErr"
	}

	fields := []struct {
		fieldPtr *string
		field    interface{}
	}{
		{&dataJson.Breadcrumb, data.Breadcrumb},
		{&dataJson.HttpData, data.HttpData},
		{&dataJson.ResourceError, data.ResourceError},
		{&dataJson.LongTask, data.LongTask},
		{&dataJson.PerformanceData, data.PerformanceData},
		{&dataJson.RequestData, data.RequestData},
		{&dataJson.Response, data.Response},
		{&dataJson.Memory, data.Memory},
		{&dataJson.CodeError, data.CodeError},
		{&dataJson.RecordScreen, data.RecordScreen},
		{&dataJson.DeviceInfo, data.DeviceInfo},
	}

	for _, f := range fields {
		if f.fieldPtr != nil {
			serialized, err := serializeField(f.field)
			if err != nil {
				// 错误处理...
				return nil
			}
			*f.fieldPtr = serialized
		}
	}

	fmt.Println("查询语句3:", dataJson)
	p := influxdb2.NewPoint(dataJson.Type,
		map[string]string{"apikeytest": "abcd"}, // tag
		untils.StructToMap(dataJson),
		time.Now()) // 时间戳
		models.WriteAPI.WritePoint(p)
		models.WriteAPI.Flush()
	return p
}

func GetMonitor(pageNum int, pageSize int, maps *models.MonitorParams) (interface{}, interface{}, error) {

	timeLayout := "2006-01-02 15:04:05"

	// 解析时间字符串
	parsedTime1, err1 := time.Parse(timeLayout, maps.StartTime)
	if err1 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, 0, err1
	}

	// 转换为时间戳
	timestamp1 := parsedTime1.Unix()

	parsedTime2, err2 := time.Parse(timeLayout, maps.EndTime)
	if err2 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, 0, err2
	}

	// 转换为时间戳
	timestamp2 := parsedTime2.Unix()

	query := fmt.Sprintf(`from(bucket:"monitor_fiber")
	|> range(start: %d, stop: %d)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
	`, // drop 丢弃不需要的字段
		timestamp1, timestamp2, maps.Type)

	if maps.Apikey != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["apikey"] == "%s")
			`, query, maps.Apikey)
	}

	if maps.UserId != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["userId"] == "%s")
			`, query, maps.UserId)
	}

	if maps.UUID != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["uuid"] == "%s")
			`, query, maps.UUID)
	}

	if maps.Name != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["name"] == "%s")
			`, query, maps.Name)
	}

	resultCount, errCount := models.QueryAPI.Query(context.Background(), query)
	if errCount != nil {
		return nil, 0, errCount
	}

	query = fmt.Sprintf(`%s
	|> drop(columns:["_start","_stop"])
	|> sort(columns: ["_time"], desc: true)
	|> limit(n: %d, offset: %d)
	`, query, pageSize, pageNum)

	// fmt.Println("查询语句:", query)

	result, err := models.QueryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	results, resErr := untils.InfluxdbQueryResult(result)

	if resErr != nil {
		return nil, 0, resErr
	}

	totalCount, errCount := untils.InfluxdbQueryResult(resultCount)

	if errCount != nil {
		return nil, 0, errCount
	}

	// 返回 JSON
	return results, len(totalCount), nil
}

func GetEchartMonitor(maps *models.MonitorParams) (interface{}, error) {

	timeLayout := "2006-01-02 15:04:05"

	// 解析时间字符串
	parsedTime1, err1 := time.Parse(timeLayout, maps.StartTime)
	if err1 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, err1
	}

	// 转换为时间戳
	timestamp1 := parsedTime1.Unix()

	parsedTime2, err2 := time.Parse(timeLayout, maps.EndTime)
	if err2 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, err2
	}

	// 转换为时间戳
	timestamp2 := parsedTime2.Unix()

	query := fmt.Sprintf(`from(bucket:"monitor_fiber")
	|> range(start: %d, stop: %d)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
	`, // drop 丢弃不需要的字段
		timestamp1, timestamp2, maps.Type)

	if maps.RecordScreenId != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["recordScreenId"] == "%s")
			`, query, maps.RecordScreenId)
	}

	if maps.Apikey != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["apikey"] == "%s")
			`, query, maps.Apikey)
	}

	if maps.Name != "" {
		query = fmt.Sprintf(`%s
			|> filter(fn: (r) => r["name"] == "%s")
			`, query, maps.Name)
	}

	query = fmt.Sprintf(`%s
	|> drop(columns:["_start","_stop"])
	`, query)

	// fmt.Println("查询语句:", query)

	result, err := models.QueryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	results, resErr := untils.InfluxdbQueryResult(result)

	if resErr != nil {
		return nil, resErr
	}

	// 返回 JSON
	return results, nil
}

func GetRecordScreen(recordScreenId string) (interface{}, error) {

	query := fmt.Sprintf(`from(bucket:"monitor_fiber")
	|> range(start: -365d)
	|> filter(fn: (r) => r["_measurement"] == "recordScreen")
	|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
	`, // drop 丢弃不需要的字段
	)

	query = fmt.Sprintf(`%s
		|> filter(fn: (r) => r["recordScreenId"] == "%s")
		`, query, recordScreenId)

	query = fmt.Sprintf(`%s
	|> drop(columns:["_start","_stop"])
	`, query)

	// fmt.Println("查询语句:", query)

	result, err := models.QueryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	results, resErr := untils.InfluxdbQueryResult(result)

	if resErr != nil {
		return nil, resErr
	}

	// 返回 JSON
	return results, nil
}
