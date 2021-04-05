package main

import (
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/controllers"
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/repository"
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/services"
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/storage"
)

type SetUpControllers struct {
	helloController *controllers.HelloController
}

func SetUpApp() *SetUpControllers {

	services := setUpServices()
	return &SetUpControllers{
		helloController: controllers.NewHelloController(services.helloService),
	}
}

type SetUpServices struct {
	helloService *services.HelloService
}

func setUpServices() *SetUpServices {

	return &SetUpServices{
		helloService: services.NewHelloService(repository.NewHelloRepository(storage.GetDB())),
	}
}
