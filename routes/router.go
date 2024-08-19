package routes

import (
	"net/http"

	"github.com/LocTr/reverse_uber_be/models"
	"github.com/LocTr/reverse_uber_be/services"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()
	initRoute(r)
	return r
}

func initRoute(r *gin.Engine) {
	_ = r.SetTrustedProxies(nil)
	r.RedirectTrailingSlash = false
	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusNotFound, c.Request.RequestURI+" not found")
	})

	r.NoMethod(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusMethodNotAllowed, c.Request.Method+" is not allowed here")
	})

}

func InitGin() {
	// why ??
	// gin.DisableConsoleColor()

	gin.SetMode(services.Config.Mode)
	// do some other initialization stuff (?)
}
