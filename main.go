package main

import (
	"wereserve/controller"
	"wereserve/initializer"
	"wereserve/middleware"
	"wereserve/repository"
	"wereserve/service"

	"github.com/gin-gonic/gin"
)

func init( ) {
	initializer.LoadEnvVariable()
	initializer.ConnectDB()
	initializer.SyncMigrate()
}

func main() {
	//? setup gin
	r := gin.Default()

	//? initialization Repository, service, controller
	//? USER
	userRepo := repository.NewUserRepository(initializer.DB)
	userService := service.NewUserService(userRepo)
	authController := controller.NewAuthController(userService)

	//? Table
	tableRepo := repository.NewTableRepository(initializer.DB)
	tableService := service.NewTableService(tableRepo)
	tableController := controller.NewTableController(tableService)

	//? Menu
	menuRepo := repository.NewMenuRepository(initializer.DB)
	menuService := service.NewMenuService(menuRepo)
	menuController := controller.NewMenuController(menuService)

	//? Reservation
	reservationRepo := repository.NewReservationRepository(initializer.DB)
	reservationService := service.NewReserveService(reservationRepo)
	reservationController := controller.NewReservationController(reservationService)


	//? middleware	
	authMiddleware := middleware.AuthMiddleware(userRepo)

	// Un-Protected route
	r.POST("/signup", authController.SignUp)
	r.POST("/login", authController.Login)

	//Table
	r.GET("/tables", tableController.GetAllTable)
	r.GET("/table/:id", tableController.GetTableByID )

	//Menu
	r.GET("/menu", menuController.GetAllMenus)


	//Protected Route
	//? Table
	r.POST("/admin/table", authMiddleware, tableController.CreateTable)
	r.PUT("/admin/table/:id", authMiddleware, tableController.UpdateTable)
	r.DELETE("/admin/table/:id", authMiddleware,tableController.DeleteTable)

	//? Menu
	r.POST("/admin/menu", authMiddleware, menuController.CreateMenu)

	//? Reserved
	r.POST("/customer/reserve", authMiddleware,reservationController.CreateReservation)
	r.GET("/customer/reserve/me", authMiddleware,reservationController.GetMyReservation)
	r.GET("/customer/reserve", authMiddleware,reservationController.GetAllReservation)

	r.Run()

}