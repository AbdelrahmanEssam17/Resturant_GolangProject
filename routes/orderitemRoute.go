package routes


 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)
func OrderItemRoute(incomingroute*gin.Engine){
incomingroute.Post("/orderitem",controllers.Createorderitem())
incomingroute.GET("/orderitem",controllers.Getorderitem())
incomingroute.Patch("/orderitem/:orderitem_id",controllers.UpdateOrderitem())
incomingroute.Get("/orderitem/:orderitem_id",controllers.Getorderitem())
}