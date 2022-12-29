package controller

type AppController struct {
	User interface {
		UserController
	}

	Book interface {
		BookController
	}

	Employee interface {
		EmployeeController
	}

	Customer interface {
		CustomerController
	}
}
