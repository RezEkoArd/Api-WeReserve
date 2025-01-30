package controller

import (
	"net/http"
	"time"
	"wereserve/models"
	"wereserve/service"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	service *service.ReserveService
}

func NewReservationController(service *service.ReserveService) *ReservationController {
	return &ReservationController{service: service}
}

type CreateReservationRequest struct {
	TableID		uint	`json:"table_id" binding:"required"`
	Date		string	`json:"date" binding:"required"`
	Time		string	`json:"time" binding:"required"`
	NumberOfPeople	int	`json:"number_of_people" binding:"required"`
}

func (c *ReservationController) CreateReservation(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	currentUser := user.(models.User)

	var req CreateReservationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// Konversi string ke Time.time (date, table data)
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invaid Date format (YYY-MM-DD)",
		})
		return
	}

	time, err := time.Parse("15:04:05", req.Time)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid time format (HH:MM:SS)",
		})
		return
	}

	reservation := models.Reservation{
		UserID: currentUser.ID,
		TableID: req.TableID,
		Date: date,
		Time: time,
		NumberOfPeople: req.NumberOfPeople,
	}

	createdReservation, err := c.service.CreateReservation(&reservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdReservation)
}

func (c *ReservationController) GetMyReservation(ctx *gin.Context){
	user, _ := ctx.Get("user")
	currentUser := user.(models.User)

	reservations, err := c.service.GetUserReservation(currentUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// Get All
func (c *ReservationController) GetAllReservation(ctx *gin.Context) {
	reservations, err := c.service.GetAllReserve()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,reservations)
}