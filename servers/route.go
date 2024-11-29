package servers

import (
	"job-application/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(h *HandlerOps) *gin.Engine {
	g := gin.New()
	g.ContextWithFallback = true
	g.Use(gin.Recovery(), middlewares.LoggerMiddleware(), middlewares.ErrorHandler)

	g.NoRoute(InvalidRoute)

	SetupAuthenRoutes(g, h)
	SetupRouteJob(g, h)

	return g
}

func SetupAuthenRoutes(g *gin.Engine, h *HandlerOps) {
	g.POST("/register", h.UserController.PostRegisterUserController)
	g.POST("/login", h.UserController.PostLoginUserController)
}

func SetupRouteJob(g *gin.Engine, h *HandlerOps) {
	g.GET("/jobs", middlewares.AuthorizationJob, h.JobController.GetAllJobListController)
	g.POST("/jobs", middlewares.AuthorizationJob, h.JobController.PostCreateJobController)
	g.POST("/jobs/:job_id", middlewares.AuthorizationJob, h.JobController.PostApplyJobController)
	g.PATCH("/jobs/:job_id", middlewares.AuthorizationJob, h.JobController.PatchCloseJobController)
}
