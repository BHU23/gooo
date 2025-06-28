package app

import "learning/app/controller"

type Application struct {
	Server *controller.Server
}

func New() *Application {
	return &Application{
		Server: controller.New(),
	}
}

func (app *Application) Run() {
	app.Server.Run()
}