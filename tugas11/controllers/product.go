package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tesjwt.go/database"
	"tesjwt.go/helpers"
	"tesjwt.go/models"
)

// CreateProduct godoc
// @Summary Create Product
// @Description Create Product for Product identified by given id
// @Tags Product
// @Accept json
// @Produce json
// @Param productId path int true "ID of the Product"
// @Param message query string true "message"
// @Security BearerAuth
// @Success 201 {object} models.Product "Create Product success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Product Not Found"
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

// UpdateProduct godoc
// @Summary Update Product
// @Description Update Product identified by given id
// @Tags Product
// @Accept json
// @Produce json
// @Param productId path int true "ID of the Product"
// @Security BearerAuth
// @Success 200 {object} models.Product "Update Product success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Product Not Found"
// @Router /products/{productId} [put]
func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete Product identified by given ID
// @Tags Product
// @Accept json
// @Produce json
// @Param ProductId path int true "ID of the Product"
// @Security BearerAuth
// @Success 200 {string} string "Delete Product success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Product Not Found"
// @Router /products/{productId} [delete]
func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}

// GetProduct godoc
// @Summary Get Product
// @Description Get Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param productId path int true "ID of the Product"
// @Security BearerAuth
// @Success 200 {object} models.Product{} "Get Product success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Product Not Found"
// @Router /products/{productId} [get]
func FindProductById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Find(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

// GetAllProduct godoc
// @Summary Get all Product
// @Description Get all existing Product
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Product{} "Get all Product success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Product Not Found"
// @Router /products [get]
func FindAllProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Product := []models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := db.Debug().Find(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}
