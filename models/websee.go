package models

import (
	"context"
	"encoding/json"
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

type LongTask struct {
	Time     int64       `json:"time"`
	Name     string      `json:"name"`
	LongTask interface{} `json:"longTask"`
}

type PerformanceData struct {
	Name   string `json:"name"`
	Value  int64  `json:"value"`
	Rating string `json:"rating"`
}

type MemoryData struct {
	Name   string `json:"name"`
	Memory struct {
		JSHeapSizeLimit int64 `json:"jsHeapSizeLimit"`
		TotalJSHeapSize int64 `json:"totalJSHeapSize"`
		UsedJSHeapSize  int64 `json:"usedJSHeapSize"`
	} `json:"memory"`
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

type ReportData struct {
	Type            string           `json:"type"`
	PageUrl         string           `json:"pageUrl"`
	Time            int64            `json:"time"`
	UUID            string           `json:"uuid"`
	Apikey          string           `json:"apikey"`
	Status          string           `json:"status"`
	SdkVersion      string           `json:"sdkVersion"`
	Breadcrumb      []*BreadcrumbData `json:"breadcrumb"`
	HttpData        *HttpData        `json:"httpData,omitempty"`
	ResourceError   *ResourceError   `json:"resourceError,omitempty"`
	LongTask        *LongTask        `json:"longTask,omitempty"`
	PerformanceData *PerformanceData `json:"performanceData,omitempty"`
	MemoryData      *MemoryData      `json:"memoryData,omitempty"`
	CodeError       *CodeError       `json:"codeError,omitempty"`
	RecordScreen    *RecordScreen    `json:"recordScreen,omitempty"`
	DeviceInfo      *DeviceInfo      `json:"deviceInfo,omitempty"`
}

type ReportDataJson struct {
	Type            string           `json:"type"`
	PageUrl         string           `json:"pageUrl"`
	Time            int64            `json:"time"`
	UUID            string           `json:"uuid"`
	Apikey          string           `json:"apikey"`
	Status          string           `json:"status"`
	SdkVersion      string           `json:"sdkVersion"`
	Breadcrumb      string `json:"breadcrumb"`
	HttpData        string        `json:"httpData,omitempty"`
	ResourceError   string   `json:"resourceError,omitempty"`
	LongTask        string        `json:"longTask,omitempty"`
	PerformanceData string `json:"performanceData,omitempty"`
	MemoryData      string      `json:"memoryData,omitempty"`
	CodeError       string       `json:"codeError,omitempty"`
	RecordScreen    string    `json:"recordScreen,omitempty"`
	DeviceInfo      string      `json:"deviceInfo,omitempty"`
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
	Type     string      `json:"type"`
	Category string      `json:"category"`
	Status   string       `json:"status"`
	Time     int64       `json:"time"`
	Data     interface{} `json:"data,omitempty"`
}

func SetMonitor(data *ReportData) *write.Point {
	dataJson := new(ReportDataJson)
	dataJson.Type = data.Type
	dataJson.PageUrl= data.PageUrl
	dataJson.Time= data.Time
	dataJson.UUID= data.UUID
	dataJson.Apikey= data.Apikey
	dataJson.Status= data.Status
	dataJson.SdkVersion= data.SdkVersion
	if (data.Breadcrumb != nil) {
		breaByt,err := json.Marshal(data.Breadcrumb)
		if err != nil {
			return nil
		}
		dataJson.Breadcrumb = string(breaByt)
	}
	if (data.HttpData != nil) {
		breaByt,err := json.Marshal(data.HttpData)
		if err != nil {
			return nil
		}
		dataJson.HttpData = string(breaByt)
	}
	if (data.ResourceError != nil) {
		breaByt,err := json.Marshal(data.ResourceError)
		if err != nil {
			return nil
		}
		dataJson.ResourceError = string(breaByt)
	}

	if (data.LongTask != nil) {
		breaByt,err := json.Marshal(data.LongTask)
		if err != nil {
			return nil
		}
		dataJson.LongTask = string(breaByt)
	}

	if (data.PerformanceData != nil) {
		breaByt,err := json.Marshal(data.PerformanceData)
		if err != nil {
			return nil
		}
		dataJson.PerformanceData = string(breaByt)
	}

	if (data.MemoryData != nil) {
		breaByt,err := json.Marshal(data.MemoryData)
		if err != nil {
			return nil
		}
		dataJson.MemoryData = string(breaByt)
	}

	if (data.CodeError != nil) {
		breaByt,err := json.Marshal(data.CodeError)
		if err != nil {
			return nil
		}
		dataJson.CodeError = string(breaByt)
	}

	if (data.RecordScreen != nil) {
		breaByt,err := json.Marshal(data.RecordScreen)
		if err != nil {
			return nil
		}
		dataJson.RecordScreen = string(breaByt)
	}

	if (data.DeviceInfo != nil) {
		breaByt,err := json.Marshal(data.DeviceInfo)
		if err != nil {
			return nil
		}
		dataJson.DeviceInfo = string(breaByt)
	}
	p := influxdb2.NewPoint(dataJson.Type,
		map[string]string{"apikeytest": "abcd"}, // tag
		untils.StructToMap(dataJson),
		time.Now()) // 时间戳
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	return p
}

func GetMonitor() (interface{}, error) {
	query := `from(bucket:"monitor_fiber")
	|> range(start: -8d)
	|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
	|> drop(columns:["_start","_stop"])` // 丢弃不需要的字段
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
