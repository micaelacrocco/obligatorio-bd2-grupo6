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

	policeAgentRepository := repositories.NewPoliceAgentRepository(database)
	policeAgentUseCase := usecases.NewPoliceAgentUseCase(policeAgentRepository)
	policeAgentHandler := handlers.NewPoliceAgentHandler(policeAgentUseCase)

	policeStationRepository := repositories.NewPoliceStationRepository(database)
	policeStationUseCase := usecases.NewPoliceStationUseCase(policeStationRepository)
	policeStationHandler := handlers.NewPoliceStationHandler(policeStationUseCase)

	tableRepository := repositories.NewTableRepository(database)
	tableUseCase := usecases.NewTableUseCase(tableRepository)
	tableHandler := handlers.NewTableHandler(tableUseCase)

	candidateRepository := repositories.NewCandidateRepository(database)
	candidateUseCase := usecases.NewCandidateUseCase(candidateRepository)
	candidateHandler := handlers.NewCandidateHandler(candidateUseCase)

	pollingPlaceRepository := repositories.NewPollingPlaceRepository(database)
	pollingPlaceUseCase := usecases.NewPollingPlaceUseCase(pollingPlaceRepository)
	pollingPlaceHandler := handlers.NewPollingPlaceHandler(pollingPlaceUseCase)

	tableMembersRepository := repositories.NewTableMemberRepository(database)
	tableMembersUseCase := usecases.NewTableMemberUseCase(tableMembersRepository)
	tableMemberHandler := handlers.NewTableMemberHandler(tableMembersUseCase)

	// Public routes
	// r.POST("/login", authHandler.Login)
	// r.POST("/register", authHandler.Register)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	r.GET("/circuits", circuitsHandler.GetAll)

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

	r.GET("/police-agents", policeAgentHandler.GetAll)
	r.GET("/police-agents/:id", policeAgentHandler.GetByCitizenID)
	r.POST("/police-agents", policeAgentHandler.Add)
	r.PUT("/police-agents", policeAgentHandler.Update)
	r.DELETE("/police-agents/:id", policeAgentHandler.Delete)

	r.GET("/police-stations", policeStationHandler.GetAll)
	r.POST("/police-stations", policeStationHandler.Add)
	r.PUT("/police-stations", policeStationHandler.Update)
	r.DELETE("/police-stations/:id", policeStationHandler.Delete)

	r.GET("/tables", tableHandler.GetAll)
	r.GET("/tables/:id", tableHandler.GetById)
	r.POST("/tables", tableHandler.Add)
	r.DELETE("/tables/:id", tableHandler.Delete)

	r.GET("/candidates", candidateHandler.GetAll)
	r.GET("/candidates/:citizen_id", candidateHandler.GetByCitizenID)
	r.POST("/candidates", candidateHandler.Add)
	r.DELETE("/candidates/:citizen_id/:list_number", candidateHandler.Delete)

	r.GET("/polling-places", pollingPlaceHandler.GetAll)
	r.GET("/polling-places/:id", pollingPlaceHandler.GetByID)
	r.POST("/polling-places", pollingPlaceHandler.Add)
	r.DELETE("/polling-places/:id", pollingPlaceHandler.Delete)

	r.PUT("/circuits", circuitsHandler.Update)
	r.DELETE("/circuits/:id", circuitsHandler.Delete)

	r.GET("/table-members", tableMemberHandler.GetAll)
	r.GET("/table-members/:citizen_id/:table_id", tableMemberHandler.GetCitizenIsTableMember)
	r.POST("/table-members", tableMemberHandler.Add)
	r.DELETE("/table-members/:citizen_id/:table_id", tableMemberHandler.Delete)

	// Admin-only routes
	protectedAdmin := protected.Group("/admin")
	protectedAdmin.Use(middlewares.RequireRoles("admin"))

	protectedAdmin.GET("/citizens", citizenHandler.GetAll)
	protectedAdmin.GET("/citizens/:ci", citizenHandler.GetById)
	protectedAdmin.POST("/citizens", citizenHandler.AddCitizen)
	protectedAdmin.PUT("/citizens/:id", citizenHandler.Update)
	protectedAdmin.DELETE("/citizens/:id", citizenHandler.Delete)

	protectedAdmin.GET("/circuits/:id", circuitsHandler.GetById)
	r.POST("/circuits", circuitsHandler.AddCircuit)

	r.Run("localhost:8080")
}
