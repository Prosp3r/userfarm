package referral

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/referral")
	// Referral routes
	routes.POST("/", h.CreateReferral)
	// User routes
	// routes.POST("/", h.CreateUser)
	// routes.GET("/:id", h.GetUser)
	// routes.PUT("/:id", h.UpdateUser)
	// routes.DELETE("/:id", h.DeleteUser)

	// routes.POST("/", h.AddUser)
	// routes.POST("/hashadduser", h.HashAddUser)

	// routes.POST("/auth", h.UserLogin)                 //Authentication
	// routes.POST("/initiateauth", h.InitiateUserLogin) //Authentication
	// routes.POST("/hashauth", h.UserLoginHash)         //Authentication
	// routes.POST("/refreshauth", h.RefreshUserLogin)   //Authentication
	// routes.POST("/destroyauth", h.DestroyUserLogin)   //Authentication
	// routes.POST("/recover", h.RecoverPass)
	// routes.POST("/verify/:vcode", h.VerifyUser)

	// secured := routes.Group("/secured").Use(middlewares.Auth())
	// secured.GET("/:id", h.GetUser)
	// secured.PUT("/:id", h.UpdateUser)
	// secured.GET("/dashboard", h.GetUserDashBoard)
	// secured.GET("/dashlinks", h.GetUserDashLinks)

}
