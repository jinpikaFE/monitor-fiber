package models

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
)

func SetMonitor(data map[string]interface{}) *write.Point {
	typeVal, ok := data["type"].(string)
	if !ok {
		typeVal = ""
	}
	p := influxdb2.NewPoint(typeVal,
		map[string]string{"apikeytest": "abcd"}, // tag
		data,
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
