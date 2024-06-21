package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products/model"
	"products/usecase"
)

type ProductController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProductList(c *gin.Context) {

	products, err := p.productUseCase.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProduct(c *gin.Context) {
	id := c.GetInt("id")

	product, err := p.productUseCase.GetProductByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) CreateProduct(c *gin.Context) {

	var product model.Product
	err := c.BindJSON(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	createdProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, createdProduct)
}
