package routes
 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)
func InvoiceRoute(incomingroute*gin.Engine){
incomingroute.Post("/invoice",controllers.CreateInvoice())
incomingroute.GET("/invoice",controllers.GetInvoice())
incomingroute.Patch("/invoice/:invoice_id",controllers.UpdateInvoice())
incomingroute.Get("/invoice/:invoice_id",controllers.GetInvoice())
}