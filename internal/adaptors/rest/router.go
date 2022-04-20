package rest

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hello-fresh-menu-planning-service/docs"
	"hello-fresh-menu-planning-service/internal/application/service"
)

func SetupRouter() *gin.Engine {
	engineEngin := gin.Default()
	engineEngin.SetTrustedProxies(nil)
	api := engineEngin.Group("/menu-planning-service/api/v1")
	{
		public := api.Group("/")
		{
			public.POST("/login", Login)
			public.POST("/signUp", SignUp)
		}
		protected := api.Group("/").Use(service.Authorize())
		{
			protected.POST("/weekly-menu-plan", CreateWeeklyMenuPlan)
			protected.GET("/weekly-menu-plan/:menuPlanId", FindWeeklyMenuPlanById)
			protected.GET("/weekly-menu-plans", FindAllWeeklyMenuPlans)
			protected.POST("/recipes", CreateRecipe)
			protected.PUT("/recipes/:recipeId", UpdateRecipe)
			protected.GET("/recipes/:recipeId", FindRecipeById)
			protected.DELETE("/recipes/:recipeId", DeleteRecipeById)
			protected.POST("/ingredients", CreateIngredient)
			protected.GET("/ingredients/:ingredientId", FindIngredientById)
			protected.DELETE("/ingredients/:ingredientId", DeleteIngredientById)
			protected.POST("/comments", CreateComment)
		}
	}
	engineEngin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engineEngin
}
