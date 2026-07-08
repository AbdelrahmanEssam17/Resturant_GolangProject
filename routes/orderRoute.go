package routes

 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)
func OrderRoute(incomingroute*gin.Engine){
incomingroute.Post("/order",controllers.CreateOrder())
incomingroute.GET("/order",controllers.GetOrder())
incomingroute.Patch("/order/:order_id",controllers.UpdateOrder())
incomingroute.Get("/order/:order_id",controllers.GetOrder())
}