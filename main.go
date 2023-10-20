package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ylfyt/meta/meta"
)

func main() {
	app := meta.New(&meta.Config{
		BaseUrl: "/api",
	})

	app.Map("GET", "/", func(c *fiber.Ctx) meta.ResponseDto {
		return meta.ResponseDto{
			Status:  http.StatusOK,
			Message: "",
			Success: true,
			Errors:  nil,
			Redirect: meta.ResponseRedirectInfo{
				Location: "https://google.com",
			},
		}
	})

	app.Run(3000)
}
