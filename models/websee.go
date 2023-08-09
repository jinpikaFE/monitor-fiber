package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
)

type HttpData struct {
	Type         *string `json:"type"`
	Method       *string `json:"method"`
	Time         int64   `json:"time"`
	Url          string  `json:"url"`          // 接口地址
	ElapsedTime  int64   `json:"elapsedTime"`  // 接口时长
	Message      string  `json:"message"`      // 接口信息
	Status       *int64  `json:"status"`       // 接口状态编码
	StatusString *string `json:"statusString"` // 接口状态
	RequestData  struct {
		HttpType string      `json:"httpType"` // 请求类型 xhr fetch
		Method   string      `json:"method"`   // 请求方式
		Data     interface{} `json:"data,omitempty"`
	} `json:"requestData,omitempty"`
	Response struct {
		Status *int64      `json:"status"` // 接口状态
		Data   interface{} `json:"data,omitempty"`
	} `json:"response,omitempty"`
}

type ResourceError struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

type Attribution struct {
	Name          string `json:"name"`
	EntryType     string `json:"entryType"`
	StartTime     int    `json:"startTime"`
	Duration      int    `json:"duration"`
	ContainerType string `json:"containerType"`
	ContainerSrc  string `json:"containerSrc"`
	ContainerID   string `json:"containerId"`
	ContainerName string `json:"containerName"`
}

type LongTask struct {
	StartTime   float32        `json:"startTime"`
	EntryType   string         `json:"entryType"`
	Duration    int            `json:"duration"`
	Name        string         `json:"name"`
	Attribution []*Attribution `json:"attribution"`
}

type PerformanceData struct {
	Name   string `json:"name"`
	Value  int64  `json:"value"`
	Rating string `json:"rating"`
}

type Memory struct {
	JSHeapSizeLimit int64 `json:"jsHeapSizeLimit"`
	TotalJSHeapSize int64 `json:"totalJSHeapSize"`
	UsedJSHeapSize  int64 `json:"usedJSHeapSize"`
}

type CodeError struct {
	Column   int64  `json:"column"`
	Line     int64  `json:"line"`
	Message  string `json:"message"`
	FileName string `json:"fileName"`
}

type Behavior struct {
	Type     string      `json:"type"`
	Category interface{} `json:"category"`
	Status   int64       `json:"status"`
	Time     int64       `json:"time"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Name     *string     `json:"name,omitempty"`
}

type RecordScreen struct {
	RecordScreenId string `json:"recordScreenId"`
	Events         string `json:"events"`
}

type DeviceInfo struct {
	BrowserVersion interface{} `json:"browserVersion"` // 版本号
	Browser        string      `json:"browser"`        // Chrome
	OSVersion      interface{} `json:"osVersion"`      // 电脑系统 10
	OS             string      `json:"os"`             // 设备系统
	UA             string      `json:"ua"`             // 设备详情
	Device         string      `json:"device"`         // 设备种类描述
	DeviceType     string      `json:"device_type"`    // 设备种类，如pc
}

type RequestData struct {
	HttpType string      `json:"httpType"`
	Method   string      `json:"method"`
	Data     interface{} `json:"data"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ReportData struct {
	Name            string            `json:"name"`
	Type            string            `json:"type"`
	PageUrl         string            `json:"pageUrl"`
	Rating          string            `json:"rating" description:"性能指标 poor or good"`
	Time            int64             `json:"time"`
	UUID            string            `json:"uuid"`
	Apikey          string            `json:"apikey"`
	Status          string            `json:"status"`
	SdkVersion      string            `json:"sdkVersion"`
	Events          string            `json:"events"`
	UserId          string            `json:"userId"`
	Line            int32             `json:"line" description:"发生错误位置"`
	Column          int32             `json:"column" description:"发生错误位置"`
	Message         string            `json:"message" description:"相关信息"`
	RecordScreenId  string            `json:"recordScreenId" description:"录屏信息id"`
	FileName        string            `json:"fileName" description:"错误文件"`
	Url             string            `json:"url" description:"请求url"`
	ElapsedTime     int32             `json:"elapsedTime" description:"接口时长"` // 接口时长
	Value           float32           `json:"value,omitempty" description:"性能指标值"`
	RequestData     *RequestData      `json:"requestData,omitempty" description:"请求数据"`
	Response        *Response         `json:"response,omitempty" description:"响应数据"`
	Breadcrumb      []*BreadcrumbData `json:"breadcrumb,omitempty" description:"用户行为"`
	HttpData        *HttpData         `json:"httpData,omitempty"`
	ResourceError   *ResourceError    `json:"resourceError,omitempty"`
	LongTask        *LongTask         `json:"longTask,omitempty"`
	PerformanceData *PerformanceData  `json:"performanceData,omitempty"`
	Memory          *Memory           `json:"memory,omitempty"`
	CodeError       *CodeError        `json:"codeError,omitempty"`
	RecordScreen    *RecordScreen     `json:"recordScreen,omitempty"`
	DeviceInfo      *DeviceInfo       `json:"deviceInfo,omitempty"`
}

type ReportDataJson struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	PageUrl         string  `json:"pageUrl"`
	Rating          string  `json:"rating" description:"性能指标 poor or good"`
	Time            int64   `json:"time"`
	UUID            string  `json:"uuid"`
	Apikey          string  `json:"apikey"`
	Status          string  `json:"status"`
	SdkVersion      string  `json:"sdkVersion"`
	Events          string  `json:"events"`
	UserId          string  `json:"userId"`
	Line            int32   `json:"line" description:"发生错误位置"`
	Column          int32   `json:"column" description:"发生错误位置"`
	Message         string  `json:"message" description:"相关信息"`
	RecordScreenId  string  `json:"recordScreenId"`
	FileName        string  `json:"fileName" description:"错误文件"`
	Url             string  `json:"url" description:"请求url"`
	ElapsedTime     int32   `json:"elapsedTime" description:"接口时长"` // 接口时长
	Value           float32 `json:"value,omitempty" description:"性能指标值"`
	RequestData     string  `json:"requestData,omitempty" description:"请求数据"`
	Response        string  `json:"response,omitempty" description:"响应数据"`
	Breadcrumb      string  `json:"breadcrumb,omitempty" description:"用户行为"`
	HttpData        string  `json:"httpData,omitempty"`
	ResourceError   string  `json:"resourceError,omitempty"`
	LongTask        string  `json:"longTask,omitempty"`
	PerformanceData string  `json:"performanceData,omitempty"`
	Memory          string  `json:"memory,omitempty"`
	CodeError       string  `json:"codeError,omitempty"`
	RecordScreen    string  `json:"recordScreen,omitempty"`
	DeviceInfo      string  `json:"deviceInfo,omitempty"`
}

type ResourceTarget struct {
	Src       *string `json:"src,omitempty"`
	Href      *string `json:"href,omitempty"`
	LocalName *string `json:"localName,omitempty"`
}

type AuthInfo struct {
	Apikey     string  `json:"apiKey"`
	SdkVersion string  `json:"sdkVersion"`
	UserId     *string `json:"userId,omitempty"`
}

type BreadcrumbData struct {
	Type     string      `json:"type" description:"事件类型"`
	Category string      `json:"category" description:"用户行为类型"`
	Status   string      `json:"status" description:"行为状态"`
	Time     int64       `json:"time"`
	Data     interface{} `json:"data,omitempty"`
}

type MonitorParams struct {
	UUID           string `query:"uuid" json:"uuid" xml:"uuid" form:"uuid"`
	UserId         string `query:"userId" json:"userId" xml:"userId" form:"userId"`
	RecordScreenId string `query:"recordScreenId" json:"recordScreenId" xml:"recordScreenId" form:"recordScreenId"`
	Apikey         string `query:"apikey" json:"apikey" xml:"apikey" form:"apikey"`
	Name           string `query:"name" json:"name" xml:"name" form:"name"`
	Type           string `validate:"required" query:"type" json:"type" xml:"type" form:"type"`
	// query tag是query参数别名，json xml，form适合post
	StartTime string `validate:"required" query:"startTime" json:"startTime" xml:"startTime" form:"startTime"`
	EndTime   string `validate:"required" query:"endTime" json:"endTime" xml:"endTime" form:"endTime"`
}

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

func SetMonitor(data *ReportData) *write.Point {
	dataJson := &ReportDataJson{
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
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	return p
}

func GetMonitor(pageNum int, pageSize int, maps *MonitorParams) (interface{}, interface{}, error) {

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

	resultCount, errCount := queryAPI.Query(context.Background(), query)
	if errCount != nil {
		return nil, 0, errCount
	}

	query = fmt.Sprintf(`%s
	|> drop(columns:["_start","_stop"])
	|> sort(columns: ["_time"], desc: true)
	|> limit(n: %d, offset: %d)
	`, query, pageSize, pageNum)

	// fmt.Println("查询语句:", query)

	result, err := queryAPI.Query(context.Background(), query)
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

func GetEchartMonitor(maps *MonitorParams) (interface{}, error) {

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

	result, err := queryAPI.Query(context.Background(), query)
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

	result, err := queryAPI.Query(context.Background(), query)
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
