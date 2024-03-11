package router

import (
	"database/sql"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/cmd/server/handler"
	backgroundXproficiency "github.com/proyecto-dnd/backend/internal/backgroundXProficiency"
	"github.com/proyecto-dnd/backend/internal/campaign"
	characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	classXspell "github.com/proyecto-dnd/backend/internal/classXSpell"
	"github.com/proyecto-dnd/backend/internal/event"
	raceXproficiency "github.com/proyecto-dnd/backend/internal/raceXProficiency"
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/spell"
	"github.com/proyecto-dnd/backend/internal/user"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
	"github.com/proyecto-dnd/backend/pkg/middleware"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
	firebaseApp *firebase.App
}

func NewRouter(engine *gin.Engine, db *sql.DB, firebaseApp *firebase.App) Router {
	return &router{
		engine:      engine,
		db:          db,
		firebaseApp: firebaseApp,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildUserRoutes()
	r.buildEventRoutes()
	r.buildCampaignRoutes()
	r.buildSessionRoutes()
	r.buildUserCampaignRoutes()
	r.buildSpellRoutes()
	r.buildClassXSpellRoutes()
	r.buildRaceXProficiencyRoutes()
	r.buildBackgroundXProficiencyRoutes()
	r.buildCharacterXSpellRoutes()
	// TODO Add other builders here	and write their functions
}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildUserRoutes() {
	userFirebaseRepository := user.NewUserFirebaseRepository(r.firebaseApp)
	userFirebaseService := user.NewServiceUser(userFirebaseRepository)
	userFirebaseHandler := handler.NewUserHandler(&userFirebaseService)

	userGroup := r.routerGroup.Group("/user")
	{
		// TODO Add Middlewares if needed
		userGroup.POST("/register", userFirebaseHandler.HandlerCreate())
		userGroup.POST("/login", userFirebaseHandler.HandlerLogin())
		userGroup.GET("", middleware.VerifySessionCookie(), userFirebaseHandler.HandlerGetAll())
		userGroup.GET("/:id", middleware.VerifySessionCookie(), userFirebaseHandler.HandlerGetById())
		userGroup.PUT("/:id", userFirebaseHandler.HandlerUpdate())
		userGroup.PATCH("/:id", userFirebaseHandler.HandlerPatch())
		userGroup.DELETE("/:id", userFirebaseHandler.HandlerDelete())
	}
}

func (r *router) buildEventRoutes() {
	eventRepository := event.NewEventRepository(r.db)
	eventService := event.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(&eventService)

	eventGroup := r.routerGroup.Group("/event")
	{
		eventGroup.POST("", eventHandler.HandlerCreate())
		eventGroup.GET("", eventHandler.HandlerGetAll())
		eventGroup.GET("/:id", eventHandler.HandlerGetById())
		eventGroup.GET("/session/:id", eventHandler.HandlerGetBySessionId())
		eventGroup.GET("/character/:id", eventHandler.HandlerGetByCharacterId())
		eventGroup.PUT("/:id", eventHandler.HandlerUpdate())
		eventGroup.DELETE("/:id", eventHandler.HandlerDelete())
	}
}

func (r *router) buildCampaignRoutes() {
	campaignRepository := campaign.NewCampaignRepository(r.db)
	sessionRepository := session.NewSessionRepository(r.db)
	campaignService := campaign.NewCampaignService(campaignRepository, sessionRepository)
	campaignHandler := handler.NewCampaignHandler(&campaignService)

	campaignGroup := r.routerGroup.Group("/campaign")
	{
		campaignGroup.POST("", campaignHandler.HandlerCreate())
		campaignGroup.GET("", campaignHandler.HandlerGetAll())
		campaignGroup.GET("/:id", campaignHandler.HandlerGetById())
		campaignGroup.GET("/user/:id", campaignHandler.HandlerGetByUserId())
		campaignGroup.PUT("/:id", campaignHandler.HandlerUpdate())
		campaignGroup.DELETE("/:id", campaignHandler.HandlerDelete())
	}
}

func (r *router) buildSessionRoutes() {
	sessionRepository := session.NewSessionRepository(r.db)
	sessionService := session.NewSessionService(sessionRepository)
	sessionHandler := handler.NewSessionHandler(&sessionService)

	sessionGroup := r.routerGroup.Group("/session")
	{
		sessionGroup.POST("", sessionHandler.HandlerCreate())
		sessionGroup.GET("", sessionHandler.HandlerGetAll())
		sessionGroup.GET("/:id", sessionHandler.HandlerGetById())
		sessionGroup.GET("/campaign/:id", sessionHandler.HandlerGetByCampaignId())
		sessionGroup.PUT("/:id", sessionHandler.HandlerUpdate())
		sessionGroup.DELETE("/:id", sessionHandler.HandlerDelete())
	}
}

func (r *router) buildUserCampaignRoutes() {
	userCampaignRepository := user_campaign.NewUserCampaignRepository(r.db)
	userCampaignService := user_campaign.NewUserCampaignService(userCampaignRepository)
	userCampaignHandler := handler.NewUserCampaignHandler(&userCampaignService)

	userCampaignGroup := r.routerGroup.Group("/user_campaign")
	{
		userCampaignGroup.POST("", userCampaignHandler.HandlerCreate())
		userCampaignGroup.GET("", userCampaignHandler.HandlerGetAll())
		userCampaignGroup.GET("/:id", userCampaignHandler.HandlerGetById())
		userCampaignGroup.GET("/campaign/:id", userCampaignHandler.HandlerGetByCampaignId())
		userCampaignGroup.GET("/user/:id", userCampaignHandler.HandlerGetByUserId())
		userCampaignGroup.DELETE("/:id", userCampaignHandler.HandlerDelete())
	}
}

func (r *router) buildSpellRoutes() {
	spellRepository := spell.NewSpellRepository(r.db)
	spellService := spell.NewSpellService(spellRepository)
	spellHandler := handler.NewSpellHandler(&spellService)

	spellGroup := r.routerGroup.Group("/spell")
	{
		spellGroup.POST("", spellHandler.HandlerCreate())
		spellGroup.GET("", spellHandler.HandlergetAll())
		spellGroup.GET("/:id", spellHandler.HandlerGetById())
		spellGroup.PUT("/:id", spellHandler.HandlerUpdate())
		spellGroup.DELETE("/:id", spellHandler.HandlerDelete())
	}
}

func (r *router) buildClassXSpellRoutes() {
	classXSpellRepository := classXspell.NewClassXSpellRepository(r.db)
	classXSpellService := classXspell.NewClassXSpelService(classXSpellRepository)
	classXSpellHandler := handler.NewClassXSpellHandler(&classXSpellService)

	classXSpellGroup := r.routerGroup.Group("/classxspell")
	{
		classXSpellGroup.POST("", classXSpellHandler.HandlerCreate())
		classXSpellGroup.DELETE("", classXSpellHandler.HandlerDelete())
	}
}

func (r *router) buildRaceXProficiencyRoutes() {
	raceXProficiencyRepository := raceXproficiency.NewRaceXProficiencyRepository(r.db)
	raceXProficiencyService := raceXproficiency.NewRaceXProficiencyService(raceXProficiencyRepository)
	raceXProficiencyHandler := handler.NewRaceXProficiencyHandler(raceXProficiencyService)

	cassXProficiencyGroup := r.routerGroup.Group("/raceXproficiency")
	{
		cassXProficiencyGroup.POST("", raceXProficiencyHandler.HandlerCreate())
		cassXProficiencyGroup.DELETE("", raceXProficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildBackgroundXProficiencyRoutes() {
	backgroundXProficiencyRepository := backgroundXproficiency.NewBackgroundXProficiencyRepository(r.db)
	backgroundXProficiencyService := backgroundXproficiency.NewBackgroundXProficiencyService(backgroundXProficiencyRepository)
	backgroundXProficiencyHandler := handler.NewBackgroundXProficiencyHandler(backgroundXProficiencyService)

	backgroundXProficiencyGroup := r.routerGroup.Group("/backgroundXproficiency")
	{
		backgroundXProficiencyGroup.POST("", backgroundXProficiencyHandler.HandlerCreate())
		backgroundXProficiencyGroup.DELETE("", backgroundXProficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildCharacterXSpellRoutes() {
	characterXSpellRepository := characterXspell.NewCharacterXSpellRepository(r.db)
	characterXSpellService := characterXspell.NewCharacterXSpellService(characterXSpellRepository)
	characterXSpellHandler := handler.NewCharacterXSpellHandler(characterXSpellService)

	characterXSpellGroup := r.routerGroup.Group("/characterXspell")
	{
		characterXSpellGroup.POST("", characterXSpellHandler.HandlerCreate())
		characterXSpellGroup.DELETE("/:id", characterXSpellHandler.HandlerDelete())
	}
}
