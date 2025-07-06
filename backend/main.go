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
	// authUseCase :∑= usecases.NewAuthUseCase(userRepo, citizenRepo)
	// authHandler := handlers.NewAuthHandler(authUseCase)

	circuitsRepo := repositories.NewCircuitRepository(database)
	circuitsUseCase := usecases.NewCircuitsUseCase(circuitsRepo)
	circuitsHandler := handlers.NewCircuitsHandler(circuitsUseCase)

	politicalPartyRepo := repositories.NewPoliticalPartyRepository(database)
	politicalPartyUseCase := usecases.NewPoliticalPartyUseCase(politicalPartyRepo)
	politicalPartyHandler := handlers.NewPoliticalPartyHandler(politicalPartyUseCase)

	partyListRepository := repositories.NewPartyListRepository(database)
	partyListUseCase := usecases.NewPartyListUseCase(partyListRepository)
	partyListHandler := handlers.NewPartyListHandler(partyListUseCase)

	listVoteRepository := repositories.NewListVoteRepository(database)
	listVoteUseCase := usecases.NewListVoteUseCase(listVoteRepository)
	listVoteHandler := handlers.NewListVoteHandler(listVoteUseCase)

	departmentRepository := repositories.NewDepartmentRepository(database)
	departmentUseCase := usecases.NewDepartmentUseCase(departmentRepository)
	departmentHandler := handlers.NewDepartmentHandler(departmentUseCase)

	zoneRepository := repositories.NewZoneRepository(database)
	zoneUseCase := usecases.NewZoneUseCase(zoneRepository)
	zoneHandler := handlers.NewZoneHandler(zoneUseCase)

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

	r.GET("/party-lists", partyListHandler.GetAll)
	r.POST("/party-lists", partyListHandler.Add)
	r.PUT("/party-lists", partyListHandler.Update)
	r.DELETE("/party-lists/:list_number", partyListHandler.Delete)

	r.GET("/list-votes", listVoteHandler.GetAll)
	r.POST("/list-votes", listVoteHandler.Add)
	r.PUT("/list-votes", listVoteHandler.Update)
	r.DELETE("/list-votes/:id", listVoteHandler.Delete)

	r.GET("/departments", departmentHandler.GetAll)
	r.POST("/departments", departmentHandler.Add)
	r.DELETE("/departments/:id", departmentHandler.Delete)

	r.GET("/zones", zoneHandler.GetAll)
	r.GET("/zones/:id", zoneHandler.GetById)
	r.POST("/zones", zoneHandler.Add)
	r.DELETE("/zones/:id", zoneHandler.Delete)

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
