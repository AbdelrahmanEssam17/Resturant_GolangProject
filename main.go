//go mod tidy download the library

// we need to import os there are a methed inside it Getenv it ti read env variables


package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/controllers"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/database"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/helper"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/middleware"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/models"
	"github.com/AbdelrahmanEssam17/Resturant_GolangProject/routes"
	" go.mongodb.org/mongo-driver/v2/mongo"
)
var foodCollection *mongo.Collection=database.OpenCollection(database.Client,"food")
func main(){
port:=os.Getenv("PORT")
if port == ""{
	port="8000"
}
//import express from 'express to generate
router:=gin.new()

router.Use(gin.Logger())
routes.Use(middleware.Authentication())

routes.UserRoute(router)
routes.FoodRoute(router)
routes.InvoiceRoute(router)
routes.MenuRoute(router)
routes.OrderRoute(router)
routes.OrderItemRoute(router)
routes.TableRoute(router)


router.Run(":"+port)
}