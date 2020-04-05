package main

import (
	"fmt"
	"github.com/coupa/foundation-go/examples/hexagon_architecture/models"
	"github.com/coupa/foundation-go/persistence"
	"log"
	"reflect"
)

func main() {
	pm, err := persistence.NewPersistenceManagerMySql("root:@/hex_demo", "cars", reflect.TypeOf(models.Car{}))
	if err != nil {
		log.Fatal(err)
	}

	findmany(pm)
	updateone(pm)
	deleteone(pm)
	createone(pm)
	findmanyload(pm)
}

func deleteone(pm persistence.PersistenceService) {
	_, err := pm.DeleteOne("13")
	if err != nil {
		log.Fatal(err)
	}
}

func updateone(pm persistence.PersistenceService) {
	car := models.Car{}
	car.Make = "daabb"
	car.LicensePlate = "123"
	car.Model = "x"
	car.Year = 22
	_, err := pm.UpdateOne("12", &car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(car)
}

func createone(pm persistence.PersistenceService) {
	car := models.Car{
		LicensePlate: "DDG 459",
		Make:         "Volvo",
		Model:        "SL",
		Year:         1985,
	}

	err := pm.CreateOne(&car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(car)
}

func findmany(pm persistence.PersistenceService) {
	params := persistence.QueryParams{
		Operands: []persistence.QueryExpression{
			{ColumnName: "year", Operator: persistence.QUERY_OPERATOR_GT, Value: "1900"},
		},
		Limit:  0,
		Offset: 0,
		Order:  []persistence.OrderStatement{{ColumnName: "id", Direction: persistence.ORDER_DIRECTION_ASC}},
	}
	cars, err := pm.FindMany(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cars)
}
func findmanyload(pm persistence.PersistenceService) {
	params := persistence.QueryParams{}
	var cars []models.Car
	err := pm.FindManyLoad(params, &cars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cars)
}
