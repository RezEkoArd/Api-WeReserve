package controller

import (
	"net/http"
	"wereserve/service"

	"github.com/gin-gonic/gin"
)


type MenuController struct {
	MenuService *service.MenuService
}

func NewMenuController(menuService *service.MenuService) *MenuController{
	return &MenuController{MenuService: menuService}
}


func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var body struct {
		Name 			string	`json:"name"` 
		Description		string	`json:"description"`
		Price			string	`json:"price"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Failed to get body",
		})
		return
	}
	
	// call Servuce
	menu, err := c.MenuService.CreateMenu(
		body.Name,body.Description,body.Price,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	//response
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Menu Created Successfully",
		"data" : gin.H{
			"name": menu.Name,
			"description": menu.Description,
			"price": menu.Price,
		},
	})
}


func (c *MenuController) GetAllMenus(ctx *gin.Context) {
	menus, err := c.MenuService.GetAllMenu()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}


