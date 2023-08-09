package models

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
)

// influxdb 的数据库实例
var influxdbClient influxdb2.Client
var org, bucket string

// 写入数据的api 实例
var WriteAPI api.WriteAPI
var QueryAPI api.QueryAPI

func init() {
	sec, err := setting.Cfg.GetSection("influxdb")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'influxdb': %v", err)
	}

	host := sec.Key("HOST").String()
	token := sec.Key("TOKEN").String()
	org = sec.Key("ORG").String()
	bucket = sec.Key("BUCKET").String()
	// Create a new InfluxDB client instance
	influxdbClient = influxdb2.NewClientWithOptions(host, token, influxdb2.DefaultOptions().SetBatchSize(100))

	// Create a new write API instance
	WriteAPI = influxdbClient.WriteAPI(org, bucket)
	QueryAPI = influxdbClient.QueryAPI(org)
}
