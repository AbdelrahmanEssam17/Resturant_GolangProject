package routes

import(
		"github.com/gin-gonic/gin"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"

)
func UserRoute(incomingroute*gin.Engine){
incomingroute.Get("/users/",controllers.GetUsers())
incomingroute.Get("/users/:user_id",controllers.GetUser())
incomingroute.Post("/users/signup",controllers.SignUp())
incomingroute.Post("/users/login",controllers.Login())
}
