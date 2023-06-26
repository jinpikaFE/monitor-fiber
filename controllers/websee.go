package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
)

// 添加监控数据
// @Summary 添加监控数据
// @Description 添加监控数据
// @Tags 添加监控数据
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
	p := models.SetMonitor(data)
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}

// 获取监控数据
// @Summary 获取监控数据
// @Description 获取监控数据
// @Tags 获取监控数据
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/monitor [get]
func GetMonitor(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	p, err := models.GetMonitor()
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}
