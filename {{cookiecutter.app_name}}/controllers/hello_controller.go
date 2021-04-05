// THIS IS A EXAMPLE FILE

package controllers

import (
	"net/http"

	services "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/services"
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/utils"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	service *services.HelloService
}

// NewMerchantController init controller by injecting services
func NewHelloController(service *services.HelloService) *HelloController {
	return &HelloController{
		service: service,
	}
}

func (controller *HelloController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": controller.service.SayHello(),
	})
}

func (controller *HelloController) GetAll(c *gin.Context) {

	data, err := controller.service.GetAll()

	if err != nil {
		utils.RespondInternalSrvErrJSON(c, utils.ErrorInternal, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
