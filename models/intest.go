package models

import (
	"context"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func SetTest() *write.Point {
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 30).
		AddField("max", 45).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	return p
}

func GetTest() (interface{}, error) {
	query := `from(bucket:"monitor_fiber")|> range(start: -8d) |> filter(fn: (r) => r._measurement == "stat")`
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	// 处理查询结果
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

	// 返回 JSON
	return results,nil
}
