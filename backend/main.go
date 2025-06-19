package main

import (
	"EleccionesUcu/db"
	"EleccionesUcu/domains/repositories"
	"EleccionesUcu/domains/usecases"
	"EleccionesUcu/handlers"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := gin.Default()
	database := db.ConnectDb()
	defer database.Close()

	circuitsRepo := repositories.NewCircuitRepository(database)
	circuitsUseCase := usecases.NewCircuitsUseCase(circuitsRepo)
	circuitsHandler := handlers.NewCircuitsHandler(circuitsUseCase)

	r.GET("/circuits/all", circuitsHandler.GetAll)
	r.Run("localhost:8080")
}
