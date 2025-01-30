package controller

import (
	"net/http"
	"strconv"
	"wereserve/models"
	"wereserve/service"

	"github.com/gin-gonic/gin"
)


type TableController struct {
	TableService *service.TableService
}

func NewTableController(tableService *service.TableService) *TableController{
	return &TableController{TableService: tableService}
}

func (c *TableController) CreateTable(ctx *gin.Context) {
	var body struct {
		TableNumber string `json:"table_number"`
		Capacity    int		`json:"capacity"`
		Status      string	`json:"status"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Failed to get body",
		})
		return
	}

	// Jika  Status Kosong maka gunakan default 
	if body.Status == "" {
		body.Status = "available"
	}

	// Call Service
	table, err := c.TableService.CreateTable(
		body.TableNumber,body.Status, body.Capacity,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	//Response
	ctx.JSON(http.StatusOK, gin.H {
		"message" : "Table Created",
		"data" : gin.H{
			"id": table.ID,
			"table_number": table.TableNumber,
			"capacity": table.Capacity,
			"status": table.Status,
		},
	})
}


func (c *TableController) UpdateTable(ctx *gin.Context) {
	// Ambil Param ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid ID",
		})
		return
	}

	// bind Request body ke strucy
	var updataTable models.Table
	if err := ctx.ShouldBindJSON(&updataTable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(), 
		})
		return
	}

	//call Service
	table, err := c.TableService.UpdateStatusTable(uint(id), &updataTable)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "berhasil di ubah",
		"data" : gin.H{
			"table" : table,
		},
	})
}

func (c *TableController) DeleteTable(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid ID", 
		})
		return
	}

	if err := c.TableService.DeleteTable(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Table Deleted",
	})
}

func (c *TableController) GetAllTable(ctx *gin.Context) {
	tables, err := c.TableService.FindAllTables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tables)
}

func (c *TableController) GetTableByID(ctx *gin.Context) {
	//Get Param ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "Get ID is Failed",
		})
		return
	}

	table, err := c.TableService.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : "table not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, table)

}