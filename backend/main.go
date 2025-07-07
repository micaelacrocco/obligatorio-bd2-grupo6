package main

import (
	"EleccionesUcu/db"
	"EleccionesUcu/domains/repositories"
	"EleccionesUcu/domains/usecases"
	"EleccionesUcu/handlers"
	"EleccionesUcu/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	// === CORS Middleware ===
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database := db.ConnectDb()
	defer database.Close()

	// Repositorios
	citizenRepo := repositories.NewCitizenRepository(database)
	userRepo := repositories.NewUserRepository(database)
	circuitsRepo := repositories.NewCircuitRepository(database)
	politicalPartyRepo := repositories.NewPoliticalPartyRepository(database)
	partyListRepository := repositories.NewPartyListRepository(database)
	listVoteRepository := repositories.NewListVoteRepository(database)
	departmentRepository := repositories.NewDepartmentRepository(database)
	zoneRepository := repositories.NewZoneRepository(database)
	policeAgentRepository := repositories.NewPoliceAgentRepository(database)
	policeStationRepository := repositories.NewPoliceStationRepository(database)
	tableRepository := repositories.NewTableRepository(database)
	candidateRepository := repositories.NewCandidateRepository(database)
	pollingPlaceRepository := repositories.NewPollingPlaceRepository(database)
	tableMembersRepository := repositories.NewTableMemberRepository(database)

	// UseCases
	citizenUseCase := usecases.NewCitizenUseCase(citizenRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)
	circuitsUseCase := usecases.NewCircuitsUseCase(circuitsRepo)
	politicalPartyUseCase := usecases.NewPoliticalPartyUseCase(politicalPartyRepo)
	partyListUseCase := usecases.NewPartyListUseCase(partyListRepository)
	listVoteUseCase := usecases.NewListVoteUseCase(listVoteRepository)
	departmentUseCase := usecases.NewDepartmentUseCase(departmentRepository)
	zoneUseCase := usecases.NewZoneUseCase(zoneRepository)
	policeAgentUseCase := usecases.NewPoliceAgentUseCase(policeAgentRepository)
	policeStationUseCase := usecases.NewPoliceStationUseCase(policeStationRepository)
	tableUseCase := usecases.NewTableUseCase(tableRepository)
	candidateUseCase := usecases.NewCandidateUseCase(candidateRepository)
	pollingPlaceUseCase := usecases.NewPollingPlaceUseCase(pollingPlaceRepository)
	tableMembersUseCase := usecases.NewTableMemberUseCase(tableMembersRepository)

	// Handlers
	citizenHandler := handlers.NewCitizenHandler(citizenUseCase)
	loginHandler := handlers.NewLoginHandler(citizenUseCase, userUseCase)
	circuitsHandler := handlers.NewCircuitsHandler(circuitsUseCase)
	politicalPartyHandler := handlers.NewPoliticalPartyHandler(politicalPartyUseCase)
	partyListHandler := handlers.NewPartyListHandler(partyListUseCase)
	listVoteHandler := handlers.NewListVoteHandler(listVoteUseCase)
	departmentHandler := handlers.NewDepartmentHandler(departmentUseCase)
	zoneHandler := handlers.NewZoneHandler(zoneUseCase)
	policeAgentHandler := handlers.NewPoliceAgentHandler(policeAgentUseCase)
	policeStationHandler := handlers.NewPoliceStationHandler(policeStationUseCase)
	tableHandler := handlers.NewTableHandler(tableUseCase)
	candidateHandler := handlers.NewCandidateHandler(candidateUseCase)
	pollingPlaceHandler := handlers.NewPollingPlaceHandler(pollingPlaceUseCase)
	tableMemberHandler := handlers.NewTableMemberHandler(tableMembersUseCase)

	// Rutas públicas
	r.POST("/auth/login", loginHandler.Login)
	r.GET("/citizens/:id", citizenHandler.GetById)
	r.GET("/party-lists", partyListHandler.GetAll)
	r.GET("/political-parties", politicalPartyHandler.GetAll)
	r.GET("/my-circuit", circuitsHandler.GetMyCircuit)

	// Rutas protegidas con middleware de autenticación
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

	//protectedAdmin.GET("/political-parties", politicalPartyHandler.GetAll)
	protectedAdmin.POST("/political-parties", politicalPartyHandler.Add)
	protectedAdmin.PUT("/political-parties", politicalPartyHandler.Update)
	protectedAdmin.DELETE("/political-parties/:id", politicalPartyHandler.Delete)

	//protectedAdmin.GET("/party-lists", partyListHandler.GetAll)
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
	//protectedAdmin.GET("/citizens/:ci", citizenHandler.GetById)
	protectedAdmin.POST("/citizens", citizenHandler.AddCitizen)
	protectedAdmin.PUT("/citizens/:id", citizenHandler.Update)
	protectedAdmin.DELETE("/citizens/:id", citizenHandler.Delete)

	r.Run(":8080") // escucha en el puerto 8080 en todas las interfaces
}
