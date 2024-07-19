package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-responsible-service/internal/service"
)

type ChildController struct {
	childservice *service.ChildService
}

func NewChildController(childservice *service.ChildService) *ChildController {
	return &ChildController{
		childservice: childservice,
	}
}

func (ct *ChildController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("vtx-responsible/api/v1")

	api.POST("/:cpf/child")
	api.GET("/:cpf/child/:rg")
	api.GET("/:cpf/child")
	api.PATCH("/:cpf/child/:rg")
	api.DELETE("/:cpf/child/:rg")

}

func (ct *ChildController) CreateChild(c *gin.Context) {

}

func (ct *ChildController) GetChild(c *gin.Context) {

}

func (ct *ChildController) FindAllChildren(c *gin.Context) {

}

func (ct *ChildController) UpdaeteChild(c *gin.Context) {

}

func (ct *ChildController) DeleteChild(c *gin.Context) {

}
