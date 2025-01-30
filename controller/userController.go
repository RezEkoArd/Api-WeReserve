package controller

import (
	"net/http"
	"wereserve/models"
	"wereserve/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

func (c *AuthController) SignUp(ctx *gin.Context){
	var body struct {
		Name 		string	`json:"name"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
		Role		models.Role	`json:"role"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to get Body",
		})
	}

	//Jika role tidak tersedia, gunakan default
	if body.Role == "" {
		body.Role = models.RoleUser
	}

	user, err := c.UserService.SignUp(body.Name, body.Email, body.Password, body.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "User Created successfully", "user" : user,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to get body",
		})
		return
	}

	// Panggil Service
	token, err := c.UserService.Login(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	//Set cookies
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*7, "","", false, true)

	//kirim response
	ctx.JSON(http.StatusOK, gin.H{
		"token" : token,
	})
}