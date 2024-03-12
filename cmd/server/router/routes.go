package router

import (
	"database/sql"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/cmd/server/handler"
	backgroundXproficiency "github.com/proyecto-dnd/backend/internal/backgroundXProficiency"
	"github.com/proyecto-dnd/backend/internal/campaign"
	characterdata "github.com/proyecto-dnd/backend/internal/characterData"
	characterXproficiency "github.com/proyecto-dnd/backend/internal/characterXProficiency"
	characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	"github.com/proyecto-dnd/backend/internal/character_feature"
	"github.com/proyecto-dnd/backend/internal/class"
	classXspell "github.com/proyecto-dnd/backend/internal/classXSpell"
	"github.com/proyecto-dnd/backend/internal/event"
	"github.com/proyecto-dnd/backend/internal/event_type"
	"github.com/proyecto-dnd/backend/internal/feature"
	"github.com/proyecto-dnd/backend/internal/friendship"
	"github.com/proyecto-dnd/backend/internal/item"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/proficiency"
	"github.com/proyecto-dnd/backend/internal/proficiencyXclass.go"
	raceXproficiency "github.com/proyecto-dnd/backend/internal/raceXProficiency"
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/skill"
	"github.com/proyecto-dnd/backend/internal/spell"
	"github.com/proyecto-dnd/backend/internal/user"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
	"github.com/proyecto-dnd/backend/internal/weapon"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.setSwaggerRoute()
	r.buildUserRoutes()
	r.buildEventRoutes()
	r.buildCampaignRoutes()
	r.buildSessionRoutes()
	r.buildClassRoutes()
	r.buildProficiencyRoutes()
	r.buildProficiencyXClassRoutes()
	r.buildUserCampaignRoutes()
	r.buildSpellRoutes()
	r.buildClassXSpellRoutes()
	r.buildRaceXProficiencyRoutes()
	r.buildBackgroundXProficiencyRoutes()
	r.buildCharacterXSpellRoutes()
	r.buildFeatureRoutes()
	r.buildItemRoutes()
	r.buildItemXCharacterDataRoutes()
	r.buildWeaponRoutes()
	r.buildWeaponXCharacterDataRoutes()
	r.buildCharacterXProficiencyRoutes()

	// TODO Add other builders here	and write their functions
}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) setSwaggerRoute() {
	r.routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		userGroup.GET("", userFirebaseHandler.HandlerGetAll())
		userGroup.GET("/:id", userFirebaseHandler.HandlerGetById())
		userGroup.PUT("/:id", userFirebaseHandler.HandlerUpdate())
		userGroup.PATCH("/:id", userFirebaseHandler.HandlerPatch())
		userGroup.DELETE("/:id", userFirebaseHandler.HandlerDelete())
	}
}

func (r *router) buildEventRoutes() {
	eventRepository := event.NewEventRepository(r.db)
	itemRepository := item.NewItemRepository(r.db)
	itemService := item.NewItemService(itemRepository)
	itemCharacterRepository := itemxcharacterdata.NewItemXCharacterDataSqlRepository(r.db)
	itemCharacterService := itemxcharacterdata.NewItemXCharacterDataService(itemCharacterRepository, itemService)
	skillrepository := skill.NewSkillRepository(r.db)
	skillService := skill.NewServiceSkill(skillrepository)
	characterRepository := characterdata.NewCharacterDataRepository(r.db)
	characterService := characterdata.NewServiceCharacterData(characterRepository, itemCharacterService, skillService)
	eventService := event.NewEventService(eventRepository, characterService)
	eventHandler := handler.NewEventHandler(&eventService)

	eventGroup := r.routerGroup.Group("/event")
	{
		eventGroup.POST("", eventHandler.HandlerCreate())
		eventGroup.GET("", eventHandler.HandlerGetAll())
		eventGroup.GET("/:id", eventHandler.HandlerGetById())
		eventGroup.GET("/type/:id", eventHandler.HandlerGetByTypeId())
		eventGroup.GET("/session/:id", eventHandler.HandlerGetBySessionId())
		eventGroup.GET("/protagonist/:id", eventHandler.HandlerGetByProtagonistId())
		eventGroup.PUT("/:id", eventHandler.HandlerUpdate())
		eventGroup.DELETE("/:id", eventHandler.HandlerDelete())
	}
}

func (r *router) buildCampaignRoutes() {
	sessionRepository := session.NewSessionRepository(r.db)
	sessionService := session.NewSessionService(sessionRepository)
	campaignRepository := campaign.NewCampaignRepository(r.db)
	campaignService := campaign.NewCampaignService(campaignRepository, sessionService)
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

func (r *router) buildClassRoutes() {
	classRepository := class.NewClassRepository(r.db)
	classService := class.NewClassService(classRepository)
	classHandler := handler.NewClassHandler(&classService)

	classGroup := r.routerGroup.Group("/class")
	{
		classGroup.POST("", classHandler.HandlerCreate())
		classGroup.GET("", classHandler.HandlerGetAll())
		classGroup.GET("/:id", classHandler.HandlerGetById())
		classGroup.PUT("/:id", classHandler.HandlerUpdate())
		classGroup.DELETE("/:id", classHandler.HandlerDelete())
	}
}

func (r *router) buildProficiencyRoutes() {
	proficiencyRepository := proficiency.NewProficiencyRepository(r.db)
	proficiencyService := proficiency.NewProficiencyService(proficiencyRepository)
	proficiencyHandler := handler.NewProficiencyHandler(&proficiencyService)

	proficiencyGroup := r.routerGroup.Group("/proficiency")
	{
		proficiencyGroup.POST("", proficiencyHandler.HandlerCreate())
		proficiencyGroup.GET("", proficiencyHandler.HandlerGetAll())
		proficiencyGroup.GET("/:id", proficiencyHandler.HandlerGetById())
		proficiencyGroup.PUT("/:id", proficiencyHandler.HandlerUpdate())
		proficiencyGroup.DELETE("/:id", proficiencyHandler.HandlerDelete())
	}
}

func (r *router) buildProficiencyXClassRoutes() {
	proficiencyXClassRepository := proficiencyXclass.NewProficiencyXClassRepository(r.db)
	proficiencyXClassService := proficiencyXclass.NewProficiencyXClassService(proficiencyXClassRepository)
	proficiencyXClassHandler := handler.NewProficiencyXClassHandler(proficiencyXClassService)

	proficiencyXClassGroup := r.routerGroup.Group("/proficiencyxclass")
	{
		proficiencyXClassGroup.POST("", proficiencyXClassHandler.HandlerCreate())
		proficiencyXClassGroup.DELETE("", proficiencyXClassHandler.HandlerDelete())
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

func (r *router) buildFriendshipRoutes() {
	friendshipRepository := friendship.NewFriendshipRepository(r.db)
	friendshipService := friendship.NewFriendshipService(friendshipRepository)
	friendshipHandler := handler.NewFriendshipHandler(friendshipService)

	friendshipGroup := r.routerGroup.Group("/friendship")
	{
		friendshipGroup.POST("", friendshipHandler.CreateHandler())
		friendshipGroup.DELETE("", friendshipHandler.DeleteHandler())
	}
}

func (r *router) buildFeatureRoutes() {
	featureRepository := feature.NewFeatureRepository(r.db)
	featureService := feature.NewFeatureService(featureRepository)
	featureHandler := handler.NewFeatureHandler(&featureService)

	featureGroup := r.routerGroup.Group("/feature")
	{
		featureGroup.POST("", featureHandler.HandlerCreate())
		featureGroup.GET("", featureHandler.HandlerGetAll())
		featureGroup.GET("/character/:id", featureHandler.HandlerGetAllFeaturesByCharacterId())
		featureGroup.GET("/:id", featureHandler.HandlerGetById())
		featureGroup.PUT("/:id", featureHandler.HandlerUpdate())
		featureGroup.DELETE("/:id", featureHandler.HandlerDelete())
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

func (r *router) buildEventTypeRoutes() {
	eventTypeRepository := event_type.NewEventTypeRepository(r.db)
	eventTypeService := event_type.NewEventTypeService(eventTypeRepository)
	eventTypeHandler := handler.NewEventTypeHandler(&eventTypeService)

	eventTypeGroup := r.routerGroup.Group("/event_type")
	{
		eventTypeGroup.POST("", eventTypeHandler.HandlerCreate())
		eventTypeGroup.GET("", eventTypeHandler.HandlerGetAll())
		eventTypeGroup.GET("/:id", eventTypeHandler.HandlerGetById())
		eventTypeGroup.GET("/:name", eventTypeHandler.HandlerGetByName())
		eventTypeGroup.PUT("/:id", eventTypeHandler.HandlerUpdate())
		eventTypeGroup.DELETE("/:id", eventTypeHandler.HandlerDelete())
	}
}

func (r *router) buildCharacterFeatureRoutes() {
	characterFeatureRepository := character_feature.NewCharacterFeatureRepository(r.db)
	characterFeatureService := character_feature.NewCharacterFeatureService(characterFeatureRepository)
	characterFeatureHandler := handler.NewCharacterFeatureHandler(&characterFeatureService)

	characterFeatureGroup := r.routerGroup.Group("/character_feature")
	{
		characterFeatureGroup.POST("", characterFeatureHandler.HandlerCreate())
		characterFeatureGroup.GET("", characterFeatureHandler.HandlerGetAll())
		characterFeatureGroup.GET("/feature/:id", characterFeatureHandler.HandlerGetByFeatureId())
		characterFeatureGroup.GET("/character/:id", characterFeatureHandler.HandlerGetByCharacterId())
		characterFeatureGroup.DELETE("/:id", characterFeatureHandler.HandlerDelete())
	}
}

func (r *router) buildItemRoutes() {
	itemRepository := item.NewItemRepository(r.db)
	itemService := item.NewItemService(itemRepository)
	itemHandler := handler.NewItemHandler(&itemService)

	itemGroup := r.routerGroup.Group("/item")
	{
		itemGroup.POST("", itemHandler.HandlerCreate())
		itemGroup.DELETE("/:id", itemHandler.HandlerDelete())
		itemGroup.GET("", itemHandler.HandlerGetAll())
		itemGroup.GET("/generic", itemHandler.HandlerGetAllGeneric())
		itemGroup.GET("/:id", itemHandler.HandlerGetById())
		itemGroup.GET("/campaign/:id", itemHandler.HandlerGetByCampaignId())
		itemGroup.PUT("/:id", itemHandler.HandlerUpdate())
	}
}

func (r *router) buildItemXCharacterDataRoutes() {
	itemRepository := item.NewItemRepository(r.db)
	itemService := item.NewItemService(itemRepository)
	itemXCharacterDataRepository := itemxcharacterdata.NewItemXCharacterDataSqlRepository(r.db)
	itemXCharacterDataService := itemxcharacterdata.NewItemXCharacterDataService(itemXCharacterDataRepository, itemService)
	itemXCharacterDataHandler := handler.NewItemXCharacterDataHandler(&itemXCharacterDataService)

	itemXCharacterDataGroup := r.routerGroup.Group("/item_character")
	{
		itemXCharacterDataGroup.POST("", itemXCharacterDataHandler.HandlerCreate())
		itemXCharacterDataGroup.DELETE("/:id", itemXCharacterDataHandler.HandlerDelete())
		itemXCharacterDataGroup.DELETE("/character/:id", itemXCharacterDataHandler.HandlerDeleteByCharacterId())
		itemXCharacterDataGroup.GET("", itemXCharacterDataHandler.HandlerGetAll())
		itemXCharacterDataGroup.GET("/:id", itemXCharacterDataHandler.HandlerGetById())
		itemXCharacterDataGroup.GET("/character/:id", itemXCharacterDataHandler.HandlerGetByCharacterDataId())
		itemXCharacterDataGroup.PUT("/:id", itemXCharacterDataHandler.HandlerUpdate())
	}
}

func (r *router) buildWeaponRoutes() {
	weaponRepository := weapon.NewWeaponRepository(r.db)
	weaponService := weapon.NewWeaponService(weaponRepository)
	weaponHandler := handler.NewWeaponHandler(&weaponService)

	weaponGroup := r.routerGroup.Group("/weapon")
	{
		weaponGroup.POST("", weaponHandler.HandlerCreate())
		weaponGroup.GET("/generic", weaponHandler.HandlerGetAllGeneric())
		weaponGroup.DELETE("/:id", weaponHandler.HandlerDelete())
		weaponGroup.GET("", weaponHandler.HandlerGetAll())
		weaponGroup.GET("/:id", weaponHandler.HandlerGetById())
		weaponGroup.GET("/campaign/:id", weaponHandler.HandlerGetByCampaignId())
		weaponGroup.PUT("/:id", weaponHandler.HandlerUpdate())
	}
}

func (r *router) buildWeaponXCharacterDataRoutes() {
	weaponRepository := weapon.NewWeaponRepository(r.db)
	weaponService := weapon.NewWeaponService(weaponRepository)
	weaponXCharacterDataRepository := weaponxcharacterdata.NewWeaponXCharacterDataSqlRepository(r.db)
	weaponXCharacterDataService := weaponxcharacterdata.NewWeaponXCharacterDataService(weaponXCharacterDataRepository, weaponService)
	weaponXCharacterDataHandler := handler.NewWeaponXCharacterDataHandler(&weaponXCharacterDataService)

	weaponXCharacterDataGroup := r.routerGroup.Group("/weapon_character")
	{
		weaponXCharacterDataGroup.POST("", weaponXCharacterDataHandler.HandlerCreate())
		weaponXCharacterDataGroup.DELETE("/:id", weaponXCharacterDataHandler.HandlerDelete())
		weaponXCharacterDataGroup.DELETE("/character/:id", weaponXCharacterDataHandler.HandlerDeleteByCharacterDataId())
		weaponXCharacterDataGroup.GET("", weaponXCharacterDataHandler.HandlerGetAll())
		weaponXCharacterDataGroup.GET("/:id", weaponXCharacterDataHandler.HandlerGetById())
		weaponXCharacterDataGroup.GET("/character/:id", weaponXCharacterDataHandler.HandlerGetByCharacterDataId())
		weaponXCharacterDataGroup.PUT("/:id", weaponXCharacterDataHandler.HandlerUpdate())
	}
}

func (r *router) buildCharacterXProficiencyRoutes() {
	characterXProficiencyRepository := characterXproficiency.NewCharacterXProficiencyRepository(r.db)
	characterXProficiencyService := characterXproficiency.NewCharacterXProficiencyService(characterXProficiencyRepository)
	characterXProficiencyHandler := handler.NewCharacterXProficiencyHandler(characterXProficiencyService)

	characterXProficiencyGroup := r.routerGroup.Group("/character_proficiency")
	{
		characterXProficiencyGroup.POST("", characterXProficiencyHandler.HandlerCreate())
		characterXProficiencyGroup.DELETE("/:id", characterXProficiencyHandler.HandlerDelete())
	}
}
