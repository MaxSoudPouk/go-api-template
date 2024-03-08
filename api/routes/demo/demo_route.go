package demo

import (
	controller "go-api-template/api/controller/demo"

	"github.com/gofiber/fiber/v2"
)

func RouteDemo(router fiber.Router) {
	utilapi := router.Group("/demo")
	// utilapioauth := utilapi.Group("/demo")
	// _ = utilapioauth

	utilapi.Post("/demo", controller.DemoCtrl)
}
