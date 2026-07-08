
package routes
 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)
func MenuRoute(incomingroute*gin.Engine){
incomingroute.Post("/menu",controllers.CreateMenu())
incomingroute.GET("/menu",controllers.GetMenu())
incomingroute.Patch("/menu/:menu_id",controllers.UpdateMenu())
incomingroute.Get("/menu/:menu_id",controllers.GetMenu())
}