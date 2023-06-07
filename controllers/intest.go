package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
)

func SetTest(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	p := models.SetTest()
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}
