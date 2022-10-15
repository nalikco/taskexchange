package handler

import (
	"errors"
	"fmt"
	"os"
	"taskexchange/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	if os.Getenv("DEV") == "true" {
		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Authorization", "content-type"},
			ExposeHeaders:   []string{"Content-Length"},
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		}))
	}

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn)
			auth.POST("/sign-up", h.signUp)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.userIdentity, h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.GET("/me", h.userIdentity, h.getMyUser)
			users.PUT("/:id", h.userIdentity, h.updateUser)
			users.DELETE("/:id", h.userIdentity, h.deleteUser)
		}

		options := api.Group("/options")
		{
			options.POST("/", h.userIdentity, h.createOption)
			options.GET("/", h.getAllOptions)
			options.GET("/:id", h.getOptionById)
			options.PUT("/:id", h.userIdentity, h.updateOption)
			options.DELETE("/:id", h.userIdentity, h.deleteOption)
		}

		events := api.Group("/events")
		{
			events.GET("/", h.userIdentity, h.findAllEvents)
			events.GET("/polling", h.userIdentity, h.pollingEvents)
			events.GET("/new", h.userIdentity, h.findNewEvents)
			events.PUT("/view-all", h.userIdentity, h.viewAllEvents)
			events.PUT("/:id", h.userIdentity, h.viewEvent)
			events.DELETE("/:id", h.userIdentity, h.deleteEvent)
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.userIdentity, h.createTask)
			tasks.POST("/excel", h.userIdentity, h.createTaskFromExcelFile)
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTaskById)
			tasks.GET("/user/:user_id", h.getUserAllTasks)
			tasks.PUT("/:id", h.userIdentity, h.updateTask)
			tasks.DELETE("/:id", h.userIdentity, h.deleteTask)
		}

		offers := api.Group("/offers")
		{
			offers.GET("/performer", h.userIdentity, h.GetPerformerActiveOffers)
			offers.POST("/", h.userIdentity, h.CreateOffer)
			offers.PUT("/:id", h.userIdentity, h.UpdateOffer)
		}

		orders := api.Group("/orders")
		{
			orders.GET("/user", h.userIdentity, h.getAllUserOrders)
			orders.GET("/performer-active", h.userIdentity, h.getAllPerformerActiveOrders)
			orders.GET("/:id", h.userIdentity, h.getOrderById)
			orders.PUT("/:id", h.userIdentity, h.updateOrder)
		}
	}

	router.NoRoute(h.serveVue)

	return router
}

func (h *Handler) serveVue(c *gin.Context) {
	scriptPath, err := os.Getwd()
	if err != nil {
		c.File("./client/dist/index.html")
	}
	filePath := fmt.Sprintf("%s/client/dist%s", scriptPath, c.Request.URL.Path)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		c.File("./client/dist/index.html")
	} else {
		c.File(filePath)
	}
}
