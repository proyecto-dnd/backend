package router

import (
	"database/sql"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/cmd/server/handler"
	"github.com/proyecto-dnd/backend/internal/event"
	"github.com/proyecto-dnd/backend/internal/campaign"
	"github.com/proyecto-dnd/backend/internal/user"
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
	// TODO Add other builders here	and write their functions
}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildUserRoutes() {
	userFirebaseRepository := user.NewUserFirebaseRepository(r.firebaseApp)
	userFirebaseService := user.NewServiceUser(userFirebaseRepository)
	userFirebaseHandler := handler.NewUserHandler(&userFirebaseService)

	// userSqlRepository := user.NewUserSqlRepository(r.db)
	// userService := user.NewServiceUser(userSqlRepository)
	// userHandler := handler.NewUserHandler(&userService)

	userGroup := r.routerGroup.Group("/user")
	{
		// TODO Add Middlewares if needed
		userGroup.POST("", userFirebaseHandler.HandlerCreate())
		userGroup.GET("", userFirebaseHandler.HandlerGetAll())
		userGroup.GET("/:id", userFirebaseHandler.HandlerGetById())
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
	campaignService := campaign.NewCampaignService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(&campaignService)

	campaignGroup := r.routerGroup.Group("/campaign")
	{
		campaignGroup.POST("", campaignHandler.HandlerCreate())
		campaignGroup.GET("", campaignHandler.HandlerGetAll())
		campaignGroup.GET("/:id", campaignHandler.HandlerGetById())
		campaignGroup.PUT("/:id", campaignHandler.HandlerUpdate())
		campaignGroup.DELETE("/:id", campaignHandler.HandlerDelete())
	}
}
