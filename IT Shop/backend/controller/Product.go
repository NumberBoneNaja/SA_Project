package controller

import (
	"main/config"
	"main/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /products
func ListProducts(c *gin.Context) {
	var products []entity.Product

	db := config.DB()

	db.Preload("Category.Owner").Preload("Brand").Find(&products)

	c.JSON(http.StatusOK, &products)
}

// GET /product/:id
func GetProductByID(c *gin.Context) {
	ID := c.Param("id")
	var product entity.Product

	db := config.DB()
	results := db.Preload("Category.Owner").Preload("Brand").First(&product, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if product.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, product)
}

// PATCH /product
func UpdateProduct(c *gin.Context) {
	ID := c.Param("id")

	var product entity.Product

	db := config.DB()
	result := db.First(&product, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}