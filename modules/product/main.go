package product

import (
	"harrysbaraini/monoedit/models"
	"harrysbaraini/monoedit/system"
	"harrysbaraini/monoedit/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Service ...
var Service = utils.Service{
	Name:   "Product",
	Prefix: "/products",
	Routes: []utils.Route{
		utils.Route{Method: "GET", Url: "/", Handler: list},
		utils.Route{Method: "GET", Url: "/:id", Handler: get},
		utils.Route{Method: "POST", Url: "/", Handler: create},
		utils.Route{Method: "PUT", Url: "/:id", Handler: update},
		utils.Route{Method: "DELETE", Url: "/:id", Handler: delete},
	},
}

type createRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
}

// functions
func list(ctx *gin.Context) {
	var products []models.Product
	system.DB.Find(&products)

	ctx.JSON(http.StatusOK, gin.H{"result": &products})
}

func get(ctx *gin.Context) {
	var product models.Product
	err := system.DB.First(&product, ctx.Param("id")).Error

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"result": &product})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"result": &err})
	}
}

func create(ctx *gin.Context) {
	var json createRequest
	if ctx.BindJSON(&json) == nil {
		product := models.Product{
			Name:        json.Name,
			Description: json.Description,
			Price:       json.Price,
		}

		err := system.DB.Create(&product).Error

		if err == nil {
			ctx.JSON(http.StatusCreated, gin.H{"result": &product})
		} else {
			ctx.JSON(http.StatusNotAcceptable, gin.H{"result": &err})
		}
	}
}

func update(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

func delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}
