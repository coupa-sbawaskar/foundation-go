package main

import (
	"github.com/coupa/foundation-go/examples/hexagon_architecture/models"
	"github.com/coupa/foundation-go/persistence"
	"github.com/coupa/foundation-go/rest"
	"github.com/coupa/foundation-go/server"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func main() {
	svr := server.Server{
		Engine: gin.New(),
	}
	//svr.UseMiddleware(middleware.RequestLogger(false))

	//Register routes

	carPersistenceManager, err := persistence.NewPersistenceManagerMySql("root:@/hex_demo", "cars", reflect.TypeOf(models.Car{}))
	if err != nil {
		log.Fatal(err)
	}
	carController := rest.NewCrudController(carPersistenceManager, nil)
	svr.Engine.GET("/cars", carController.FindMany)
	svr.Engine.GET("/cars/:id", carController.FindOne)
	svr.Engine.POST("/cars", carController.CreateOne)
	svr.Engine.PUT("/cars/:id", carController.UpdateOne)
	svr.Engine.DELETE("/cars/:id", carController.DeleteOne)

	svr.Engine.Run(":8080") //svr.Engine.Run() without address parameter will run on ":8080"

}
