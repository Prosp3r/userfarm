package main

import (
	"github.com/Prosp3r/deepseek/userfarm/pkg/common/db"
	"github.com/Prosp3r/deepseek/userfarm/pkg/referral"
	"github.com/Prosp3r/deepseek/userfarm/pkg/task"
	"github.com/Prosp3r/deepseek/userfarm/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.Use(CORSMiddleware())

	logic.LogicDB(h) //Initiate db connections for logic con different from routs
	go logic.Cacheurl()

	// urls.RegisterRoutes(r, h)
	user.RegisterRoutes(r, h)
	task.RegisterRoutes(r, h)
	referral.RegisterRoutes(r, h)
	// points.RegisterRoutes(r, h)
	// subscriptions.RegisterRoutes(r, h)
	// tokenbank.RegisterRoutes(r, h)
	// tokenlist.RegisterRoutes(r, h)
	// wager.RegisterRoutes(r, h)

	// tokenlist.RegisterRoutes(r, h)
	// wager.RegisterRoutes(r, h)

	//register more routes

	r.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
