package main

import "github.com/gofiber/fiber/v2"

type EndPoint struct {
	Method      string
	Path        string
	HandlerFunc any
	Middlewares []any
}

type middleware struct {
	Paths   []string
	Handler any
}

type Config struct {
	BaseUrl string
}

type App struct {
	fiberApp     *fiber.App
	router       fiber.Router
	config       *Config
	endPoints    []EndPoint
	dependencies map[string]any
	mids         []middleware
}

type Error struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

type ResponseDto struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
	Errors  []Error `json:"errors"`
	Data    any     `json:"data"`
	Total   int64   `json:"total"`
	Count   int64   `json:"count"`
}
