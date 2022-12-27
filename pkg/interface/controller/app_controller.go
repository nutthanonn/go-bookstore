package controller

type AppController struct {
	User interface {
		UserController
	}
	Book interface {
		BookController
	}
}
