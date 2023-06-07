package models

import (
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
