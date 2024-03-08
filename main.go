package main

import (
	"fmt"
	"go-api-template/api/routes"
	"go-api-template/bootstrap"

	"log"
)

func main() {

	app := bootstrap.App()
	globalEnv := app.Env
	fiber := app.Fiber

	routes.Setup(fiber)

	log.Fatal(fiber.Listen(fmt.Sprintf(":%v", globalEnv.App.Port)))
}
