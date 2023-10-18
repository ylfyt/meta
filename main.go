package main

import (
	"net/http"

	"github.com/ylfyt/meta/meta"
)

func main() {
	app := meta.New(&meta.Config{
		BaseUrl: "/api",
	})

	app.Map("GET", "/", func() meta.ResponseDto {
		return meta.ResponseDto{
			Status:  http.StatusOK,
			Message: "",
			Success: true,
			Errors:  nil,
		}
	})

	app.Run(3000)
}
