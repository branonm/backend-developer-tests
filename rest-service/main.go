package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/services"

)

func main() {

	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	e := echo.New()
	e.GET("/people", services.GetPeople)
	e.GET("/people/:id", services.GetPersonByID)
	e.Logger.Fatal(e.Start(":8888"))

}
