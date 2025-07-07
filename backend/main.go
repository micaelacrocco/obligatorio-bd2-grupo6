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

	loginHandler := handlers.NewLoginHandler(citizenUseCase, circuitsUseCase)

	r.POST("/login", loginHandler.Login)

	// Protected admin routes
	protectedAdmin := r.Group("/")
	protectedAdmin.Use(middlewares.AuthMiddleware())

	protectedAdmin.GET("/circuits", circuitsHandler.GetAll)
	protectedAdmin.GET("/circuits/:id", circuitsHandler.GetById)
	protectedAdmin.POST("/circuits", circuitsHandler.AddCircuit)
	protectedAdmin.PUT("/circuits", circuitsHandler.Update)
	protectedAdmin.DELETE("/circuits/:id", circuitsHandler.Delete)
	protectedAdmin.GET("/circuits/:id/votes-by-party", circuitsHandler.GetVotesByParty)
	protectedAdmin.GET("/circuits/:id/results", circuitsHandler.GetVotes)
	protectedAdmin.GET("/circuits/:id/candidates/results", circuitsHandler.GetVotesByAllCandidates)

	protectedAdmin.GET("/political-parties", politicalPartyHandler.GetAll)
	protectedAdmin.POST("/political-parties", politicalPartyHandler.Add)
	protectedAdmin.PUT("/political-parties", politicalPartyHandler.Update)
	protectedAdmin.DELETE("/political-parties/:id", politicalPartyHandler.Delete)

	protectedAdmin.GET("/party-lists", partyListHandler.GetAll)
	protectedAdmin.POST("/party-lists", partyListHandler.Add)
	protectedAdmin.PUT("/party-lists", partyListHandler.Update)
	protectedAdmin.DELETE("/party-lists/:list_number", partyListHandler.Delete)

	protectedAdmin.GET("/list-votes", listVoteHandler.GetAll)
	protectedAdmin.POST("/list-votes", listVoteHandler.Add)
	protectedAdmin.PUT("/list-votes", listVoteHandler.Update)
	protectedAdmin.DELETE("/list-votes/:id", listVoteHandler.Delete)

	protectedAdmin.GET("/departments", departmentHandler.GetAll)
	protectedAdmin.POST("/departments", departmentHandler.Add)
	protectedAdmin.DELETE("/departments/:id", departmentHandler.Delete)

	protectedAdmin.GET("/zones", zoneHandler.GetAll)
	protectedAdmin.GET("/zones/:id", zoneHandler.GetById)
	protectedAdmin.POST("/zones", zoneHandler.Add)
	protectedAdmin.DELETE("/zones/:id", zoneHandler.Delete)

	protectedAdmin.GET("/police-agents", policeAgentHandler.GetAll)
	protectedAdmin.GET("/police-agents/:id", policeAgentHandler.GetByCitizenID)
	protectedAdmin.POST("/police-agents", policeAgentHandler.Add)
	protectedAdmin.PUT("/police-agents", policeAgentHandler.Update)
	protectedAdmin.DELETE("/police-agents/:id", policeAgentHandler.Delete)

	protectedAdmin.GET("/police-stations", policeStationHandler.GetAll)
	protectedAdmin.POST("/police-stations", policeStationHandler.Add)
	protectedAdmin.PUT("/police-stations", policeStationHandler.Update)
	protectedAdmin.DELETE("/police-stations/:id", policeStationHandler.Delete)

	protectedAdmin.GET("/tables", tableHandler.GetAll)
	protectedAdmin.GET("/tables/:id", tableHandler.GetById)
	protectedAdmin.POST("/tables", tableHandler.Add)
	protectedAdmin.DELETE("/tables/:id", tableHandler.Delete)

	protectedAdmin.GET("/candidates", candidateHandler.GetAll)
	protectedAdmin.GET("/candidates/:citizen_id", candidateHandler.GetByCitizenID)
	protectedAdmin.POST("/candidates", candidateHandler.Add)
	protectedAdmin.DELETE("/candidates/:citizen_id/:list_number", candidateHandler.Delete)

	protectedAdmin.GET("/polling-places", pollingPlaceHandler.GetAll)
	protectedAdmin.GET("/polling-places/:id", pollingPlaceHandler.GetByID)
	protectedAdmin.POST("/polling-places", pollingPlaceHandler.Add)
	protectedAdmin.DELETE("/polling-places/:id", pollingPlaceHandler.Delete)

	protectedAdmin.GET("/table-members", tableMemberHandler.GetAll)
	protectedAdmin.GET("/table-members/:citizen_id/:table_id", tableMemberHandler.GetCitizenIsTableMember)
	protectedAdmin.POST("/table-members", tableMemberHandler.Add)
	protectedAdmin.DELETE("/table-members/:citizen_id/:table_id", tableMemberHandler.Delete)

	protectedAdmin.GET("/citizens", citizenHandler.GetAll)
	protectedAdmin.GET("/citizens/:ci", citizenHandler.GetById)
	protectedAdmin.POST("/citizens", citizenHandler.AddCitizen)
	protectedAdmin.PUT("/citizens/:id", citizenHandler.Update)
	protectedAdmin.DELETE("/citizens/:id", citizenHandler.Delete)

	r.Run("localhost:8080")
}
