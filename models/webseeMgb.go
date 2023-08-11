package models

type MgbMonitorParams struct {
	MgbMonitorCondition
	MgbMonitorInequality
}

type MgbMonitorCondition struct {
	UUID           string `query:"uuid" json:"uuid" xml:"uuid" form:"uuid"`
	UserId         string `query:"userId" json:"userId" xml:"userId" form:"userId"`
	RecordScreenId string `query:"recordScreenId" json:"recordScreenId" xml:"recordScreenId" form:"recordScreenId"`
	// 项目名，项目key
	Apikey string `query:"apikey" json:"apikey" xml:"apikey" form:"apikey"`
	Name   string `query:"name" json:"name" xml:"name" form:"name"`
	// 类型对应不同数据表或者是集合
	Type string `validate:"required" query:"type" json:"type" xml:"type" form:"type"`
}

type MgbMonitorInequality struct {
	StartTime string `validate:"required" query:"startTime" json:"startTime" xml:"startTime" form:"startTime"`
	EndTime   string `validate:"required" query:"endTime" json:"endTime" xml:"endTime" form:"endTime"`
}
