package models

import (
	"context"
	"encoding/json"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
)

func SetTest() *write.Point {
	b, err := json.Marshal(map[string]interface{}{
		"browserVersion": "114.0.0.0",
		"browser":        "Chrome",
		"osVersion":      "10",
		"os":             "Windows",
		"ua":             "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"device":         "Unknow",
		"device_type":    "Pc",
	})
	if err != nil {
		return nil
	}
	p := influxdb2.NewPoint("whiteScreen",
		map[string]string{"apikey": "abcd"}, // tag
		map[string]interface{}{ // fields
			"type":       "whiteScreen",
			"time":       1686548461856,
			"status":     "ok",
			"userId":     "123",
			"sdkVersion": "4.0.2",
			"uuid":       "e7dec5f4-b603-4d5f-9921-b7b983d9800d",
			"pageUrl":    "http://localhost:8083/#/",
			"deviceInfo": string(b),
		},
		time.Now()) // 时间戳
	// write point asynchronously
	writeAPI.WritePoint(p)

	// // create point using fluent style
	// p = influxdb2.NewPointWithMeasurement("stat").
	// 	AddTag("unit", "temperature").
	// 	AddField("avg", 23.2).
	// 	AddField("max", 45).
	// 	SetTime(time.Now())
	// // write point asynchronously
	// writeAPI.WritePoint(p)
	writeAPI.Flush()
	return p
}

func GetTest() (interface{}, error) {
	query := `from(bucket:"monitor_fiber")|> range(start: -8d) |> filter(fn: (r) => r._measurement == "whiteScreen") |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
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
