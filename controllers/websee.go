package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/e"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
	"github.com/jinpikaFE/go_fiber/services"
)

// 添加监控数据
// @Summary 添加监控数据
// @Description 添加监控数据
// @Tags 监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/monitor [post]
func SetMonitor(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	data := &models.ReportData{}

	switch ct := c.Get("Content-Type"); ct {
	case "application/json":
		if err := json.Unmarshal(c.Body(), &data); err != nil {
			return err
		}
	case "text/plain;charset=UTF-8":
		if err := json.Unmarshal([]byte(c.Body()), &data); err != nil {
			return err
		}
	case "application/xml":
		if err := xml.Unmarshal(c.Body(), &data); err != nil {
			return err
		}
	default:
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", fmt.Errorf("unsupported Content-Type: %v", ct))
	}
	p := services.SetMonitor(data)
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}

// 获取监控数据
// @Summary 获取监控数据
// @Description 获取监控数据
// @Tags 监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/monitor [get]
func GetMonitor(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	page := &e.PageStruct{}
	maps := &models.MonitorParams{}
	err2 := c.QueryParser(page)

	if err := c.QueryParser(maps); err != nil && err2 != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*page); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*maps); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	p, resultCount, err := services.GetMonitor((page.Page-1)*page.PageSize, page.PageSize, maps)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	result := make(map[string]interface{})
	result["list"] = p
	result["pageNum"] = page.Page
	result["pageSize"] = page.PageSize
	result["total"] = resultCount
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", result)
}

// 获取监控图表数据
// @Summary 获取监控图表数据
// @Description 获取监控图表数据
// @Tags 监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/monitor/echart [get]
func GetEchartMonitor(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	maps := &models.MonitorParams{}
	if err := c.QueryParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*maps); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	p, err := services.GetEchartMonitor(maps)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	result := make(map[string]interface{})
	result["list"] = p
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", result)
}

// 获取录屏数据
// @Summary 获取录屏数据
// @Description 获取录屏数据
// @Tags 监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/monitor/screen/:id [get]
func GetRecordScreen(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id := c.Params("id")

	p, err := services.GetRecordScreen(id)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}
