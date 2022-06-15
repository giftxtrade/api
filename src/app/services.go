package app

import "github.com/giftxtrade/api/src/services"

func (app *AppBase) CreateServices() *AppBase {
	conn := app.DB
	app.UserService = &services.UserService{
		Service: services.New(conn, "users"),
	}
	app.CategoryService = &services.CategoryService{
		Service: services.New(conn, "categories"),
	}
	app.ProductService = &services.ProductService{
		Service: services.New(conn, "products"),
		CategoryService: app.CategoryService,
	}
	return app
}
