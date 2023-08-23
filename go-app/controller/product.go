package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/erajaya/go-app/helper"
	"github.com/marioheryanto/erajaya/go-app/library"
	"github.com/marioheryanto/erajaya/go-app/model"
)

type ProductController struct {
	Lib library.ProductLibraryInterface
}

type ProductControllerInterface interface {
	AddProduct(c *gin.Context)
	GetProduct(c *gin.Context)
}

func NewProductController(lib library.ProductLibraryInterface) ProductControllerInterface {
	return ProductController{
		Lib: lib,
	}
}

func (ctrl ProductController) AddProduct(c *gin.Context) {
	request := model.Product{}
	response := model.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = ctrl.Lib.AddProduct(context.Background(), request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Product added"
	c.JSON(http.StatusCreated, response)
}

func (ctrl ProductController) GetProduct(c *gin.Context) {
	response := model.Response{}
	sort := c.QueryArray("sort")

	if len(sort) > 1 {
		c.JSON(helper.GenerateResponse(c, &response, helper.NewServiceError(http.StatusBadRequest, "sort by hanya bisa 1 category")))
		return
	}

	products, err := ctrl.Lib.GetProduct(context.Background(), sort)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = products
	response.Message = fmt.Sprintf("success get %v products", len(products))
	c.JSON(http.StatusOK, response)
}
