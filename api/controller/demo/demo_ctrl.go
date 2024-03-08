package controller

import (
	"go-api-template/api/middleware"
	models "go-api-template/domain/models/demo"
	demoSvc "go-api-template/services/demo"

	"github.com/gofiber/fiber/v2"
)

func DemoCtrl(ctx *fiber.Ctx) error {
	body := models.DemoRequest{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	data, err := demoSvc.DemoSvc(body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, data)
}
