package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ylfyt/meta/meta"
)

type HomeController struct {
	hiMom *HiMom
}

func (me *HomeController) Setup(hiMom *HiMom) []meta.EndPoint {
	// me.hiMom = hiMom
	return []meta.EndPoint{
		{
			Method:      "GET",
			Path:        "/ping",
			HandlerFunc: me.ping,
			Middlewares: []any{
				func(next func() error) error {
					fmt.Println("pung")
					return next()
				},
			},
		},
	}
}

func (me *HomeController) ping() meta.ResponseDto {
	fmt.Println("getUser")
	me.hiMom.HiMom()
	return meta.ResponseDto{
		Status:  200,
		Success: true,
		Data:    "pong",
	}
}

type HiMom struct {
}

func (me *HiMom) HiMom() {
	fmt.Println("Hi, Mom!")
}

func main() {
	app := meta.New(&meta.Config{
		BaseUrl: "",
	})

	app.Map("GET", "/", func(c *fiber.Ctx) meta.ResponseDto {
		return meta.ResponseDto{
			Status:  http.StatusOK,
			Message: "",
			Success: true,
			Errors:  nil,
		}
	})

	hiMom := &HiMom{}
	app.AddController(&HomeController{})
	app.AddService(hiMom)

	app.Run(3000)
}
