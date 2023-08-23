package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/erajaya/go-app/controller"
)

func ProductRoutes(app *gin.Engine, c controller.ProductControllerInterface) {
	product := app.Group("/product")
	product.POST("/", c.AddProduct)
	product.GET("/", c.GetProduct)

}
