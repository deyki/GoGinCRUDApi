package server

import (
	"github.com/crudapigin/deyki/v2/controller"
	"github.com/crudapigin/deyki/v2/service"
)

func Run() {
	service.ConnectDB()
	controller.GinRouter()
}