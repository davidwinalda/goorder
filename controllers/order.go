package controllers

import (
	"goorder/config"
	"goorder/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	c.ShouldBindJSON(&order)

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrderById(c *gin.Context) {
	id := c.Params.ByName("id")

	var order models.Order
	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// func UpdateOrderById(c *gin.Context) {
// 	id := c.Params.ByName("id")

// 	var order models.Order

// 	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.ShouldBindJSON(&order)

// 	if err := config.DB.Save(&order).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, order)
// }
