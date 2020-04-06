package main

import (
	"github.com/coupa/foundation-go/examples/hexagon_architecture/models"
	"github.com/coupa/foundation-go/middleware"
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
	svr.UseMiddleware(middleware.RequestLogger(false))

	carPersistenceManager, err := persistence.NewPersistenceServiceMySql("root:@/hex_demo", "cars", reflect.TypeOf(models.Car{}))
	if err != nil {
		log.Fatal(err)
	}

	//Register routes
	carController := rest.NewCrudController(carPersistenceManager, nil)

	svr.Engine.GET("/cars", carController.Index)
	svr.Engine.GET("/cars/:id", carController.Show)
	svr.Engine.POST("/cars", carController.Create)
	svr.Engine.PUT("/cars/:id", carController.Update)
	svr.Engine.DELETE("/cars/:id", carController.Destroy)

	svr.Engine.Run(":8080") //svr.Engine.Run() without address parameter will run on ":8080"

}
