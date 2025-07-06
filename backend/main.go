package main

import (
	"EleccionesUcu/db"
	"EleccionesUcu/domains/repositories"
	"EleccionesUcu/domains/usecases"
	"EleccionesUcu/handlers"
	"EleccionesUcu/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database := db.ConnectDb()
	defer database.Close()

	// Dependency injection
	citizenRepo := repositories.NewCitizenRepository(database)
	citizenUseCase := usecases.NewCitizenUseCase(citizenRepo)
	citizenHandler := handlers.NewCitizenHandler(citizenUseCase)

	// userRepo := repositories.NewUserRepository(database)
	// authUseCase := usecases.NewAuthUseCase(userRepo, citizenRepo)
	// authHandler := handlers.NewAuthHandler(authUseCase)

	circuitsRepo := repositories.NewCircuitRepository(database)
	circuitsUseCase := usecases.NewCircuitsUseCase(circuitsRepo)
	circuitsHandler := handlers.NewCircuitsHandler(circuitsUseCase)

	politicalPartyRepo := repositories.NewPoliticalPartyRepository(database)
	politicalPartyUseCase := usecases.NewPoliticalPartyUseCase(politicalPartyRepo)
	politicalPartyHandler := handlers.NewPoliticalPartyHandler(politicalPartyUseCase)

	// Public routes
	// r.POST("/login", authHandler.Login)
	// r.POST("/register", authHandler.Register)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/circuits", circuitsHandler.GetAll)

	r.GET("/political-parties", politicalPartyHandler.GetAll)
	r.POST("/political-parties", politicalPartyHandler.Add)
	r.PUT("/political-parties", politicalPartyHandler.Update)
	r.DELETE("/political-parties/:id", politicalPartyHandler.Delete)
	// Admin-only routes
	protectedAdmin := protected.Group("/admin")
	protectedAdmin.Use(middlewares.RequireRoles("admin"))

	protectedAdmin.GET("/citizens", citizenHandler.GetAll)
	protectedAdmin.GET("/citizens/:ci", citizenHandler.GetById)
	protectedAdmin.POST("/citizens", citizenHandler.AddCitizen)
	protectedAdmin.PUT("/citizens/:id", citizenHandler.Update)
	protectedAdmin.DELETE("/citizens/:id", citizenHandler.Delete)

	protectedAdmin.GET("/circuits/:id", circuitsHandler.GetById)
	protectedAdmin.POST("/circuits", circuitsHandler.AddCircuit)

	r.Run("localhost:8080")
}
