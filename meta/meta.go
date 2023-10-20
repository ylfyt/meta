package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func New(config *Config) *App {
	fiberApp := fiber.New(fiber.Config{
		StrictRouting: false,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	router := fiberApp.Group(config.BaseUrl)

	return &App{
		fiberApp:     fiberApp,
		config:       config,
		router:       router,
		dependencies: make(map[string]any),
	}
}

func (app *App) Map(method string, path string, handler any, middlewares ...any) {
	app.endPoints = append(app.endPoints, EndPoint{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
		Middlewares: middlewares,
	})
}

func (app *App) AddEndPoint(endPoints ...EndPoint) {
	app.endPoints = append(app.endPoints, endPoints...)
}

func (app *App) Run(port int) {
	for _, mid := range app.mids {
		err := app.validateMiddleware(mid.Handler)
		if err != nil {
			fmt.Println("Data:", mid)
			panic(err)
		}
	}

	err := app.setupController()
	if err != nil {
		fmt.Println("ERROR", err)
		panic(err)
	}

	for _, v := range app.endPoints {
		for _, mid := range v.Middlewares {
			err := app.validateMiddleware(mid)
			if err != nil {
				fmt.Println("Data:", v, mid)
				panic(err)
			}
		}

		err := app.validateHandler(v.HandlerFunc)
		if err != nil {
			fmt.Print(v.Path, " ")
			panic(err)
		}
	}
	app.setup()
	fmt.Println("App running on port", port)
	app.fiberApp.Listen(fmt.Sprintf("0.0.0.0:%d", port))
}

func (app *App) AddService(service any) {
	if reflect.TypeOf(service).Kind() != reflect.Pointer {
		panic("service must a pointer")
	}
	app.dependencies[reflect.TypeOf(service).String()] = service
}

func (app *App) validateMiddleware(handler any) error {
	ref := reflect.TypeOf(handler)
	if ref.Kind() != reflect.Func {
		return errors.New("handler should be a function")
	}

	if ref.NumOut() < 1 || ref.NumOut() > 1 || ref.Out(0).String() != "error" {
		return errors.New("return type must be 'error'")
	}

	if ref.NumIn() < 1 || (ref.In(0).String() != "func() error" && ref.In(0).String() != "*fiber.Ctx") {
		return errors.New("first parameter must be 'next func() error' or 'c *fiber.Ctx'")
	}

	for i := 1; i < ref.NumIn(); i++ {
		if app.dependencies[ref.In(i).String()] != nil {
			continue
		}
		if ref.In(i).String() != "*fiber.Ctx" {
			return errors.New("arg only allowed with type *fiber.ctx, or from dependency services")
		}
	}
	return nil
}

func (app *App) Use(handler any, paths ...string) {
	app.mids = append(app.mids, middleware{
		Paths:   paths,
		Handler: handler,
	})
}
