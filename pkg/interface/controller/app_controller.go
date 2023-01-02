package controller

type AppController struct {
	Book interface {
		BookController
	}

	Employee interface {
		EmployeeController
	}

	Customer interface {
		CustomerController
	}

	Inventory interface {
		InventoryController
	}
}
