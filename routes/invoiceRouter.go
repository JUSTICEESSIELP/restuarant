package routes

import (
	controller "restaurant/controller"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", controller.UpdateInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())

}
