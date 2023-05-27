package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(app *gin.Engine) {
	app.Use(middleware.CORSMiddleware())

	auth := app.Group("/auth")
	api := app.Group("/api", middleware.AuthMiddleware())

	auth.POST("/signIn", controllers.Auth)
	auth.POST("/signUp", controllers.SignUp)

	api.GET("/users/:id", controllers.GetUserById)
	api.GET("/users", controllers.GetAllUsers)

	api.POST("/vacancies", controllers.PostVacancy)
	api.GET("/vacancies", controllers.GetAllVacancies)
	api.GET("/vacancies/:id", controllers.GetVacancyById)

	api.POST("/cv", controllers.PostCV)
	api.GET("/cv/:id", controllers.GetCVById)

	api.POST("/applies", controllers.PostApply)
	api.POST("/testTask", controllers.PostTestTask)
}
