package routes

import (
	controller "restaurant/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/order-items", controller.GetOrderItems())
	incomingRoutes.GET("/order-items/:item_id", controller.GetOrderItem())

	incomingRoutes.POST("/order-items", controller.CreateOrderItem())

	incomingRoutes.PATCH("/order-items/:order-id", controller.UpdatedOrderItem())

}
