package meta

import (
	"reflect"
)

const controller_INIT_FUNC_NAME = "Setup"

func (app *App) AddController(c any) {
	ref := reflect.TypeOf(c)
	if ref.Kind() != reflect.Pointer {
		panic("please supply pointer to controller")
	}
	method, ok := reflect.TypeOf(c).MethodByName(controller_INIT_FUNC_NAME)
	if !ok {
		panic("controller must be implement setup method")
	}
	if method.Type.NumOut() == 0 || method.Type.Out(0) != reflect.TypeOf([]EndPoint{}) {
		panic("the setup method must be return slice of endpoint")
	}
	app.controllers = append(app.controllers, c)
}
