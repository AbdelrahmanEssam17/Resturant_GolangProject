package routes
 
import(
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/gin-gonic/gin"
)


func FoodRoute(incomingroute*gin.Engine){
incomingroute.Post("/foods",controllers.CreateFoods())
incomingroute.GET("/foods",controllers.GetFood())
incomingroute.Patch("/foods/:food_id",controllers.UpdateFood())
incomingroute.Get("/foods/:food_id",controllers.GetFood())
}