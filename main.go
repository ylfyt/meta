package main

import "net/http"

func main() {
	app := New(&Config{
		BaseUrl: "/api",
	})

	app.Map("GET", "/", func() ResponseDto {
		return ResponseDto{
			Status:  http.StatusOK,
			Message: "",
			Success: true,
			Errors:  nil,
		}
	})

	app.Run(3000)
}
