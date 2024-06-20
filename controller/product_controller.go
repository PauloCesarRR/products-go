package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products/model"
)

type ProductController struct {
	//TODO: UseCase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProductList(c *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Beetle Juice",
			Price: 20.50,
		},
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProduct(c *gin.Context) {
	id := c.GetInt("id")
	product := model.Product{
		ID:    id,
		Name:  "beetle juice",
		Price: 20.50,
	}

	c.JSON(http.StatusOK, product)
}
