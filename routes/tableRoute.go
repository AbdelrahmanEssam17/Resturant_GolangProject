package routes
package routes

 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)
func TableRoute(incomingroute*gin.Engine){
incomingroute.Post("/table",controllers.CreateTable())
incomingroute.GET("/table",controllers.GetTable())
incomingroute.Patch("/table/:table_id",controllers.UpdateTable())
incomingroute.Get("/table/:table_id",controllers.GetTable())
}