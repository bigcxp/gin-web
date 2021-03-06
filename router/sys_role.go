package router

import (
	v1 "gin-web/api/v1"
	"gin-web/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 角色路由
func InitRoleRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("/role").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", v1.GetRoles)
		router.POST("/create", v1.CreateRole)
		router.PATCH("/update/:roleId", v1.UpdateRoleById)
		router.PATCH("/menus/update/:roleId", v1.UpdateRoleMenusById)
		router.PATCH("/apis/update/:roleId", v1.UpdateRoleApisById)
		router.DELETE("/delete/batch", v1.BatchDeleteRoleByIds)
	}
	return router
}
