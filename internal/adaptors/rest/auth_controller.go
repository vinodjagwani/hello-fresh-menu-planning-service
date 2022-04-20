package rest

import (
	"github.com/gin-gonic/gin"
	profile "github.com/lvornholt/go-profiles"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/ports"
	"hello-fresh-menu-planning-service/internal/application/service"
	"hello-fresh-menu-planning-service/internal/application/usecase"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"net/http"
)

// @BasePath /api/v1
// HelloFresh godoc
// @Summary User login
// @Schemes
// @Description body
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param menuPlan body dto.LoginRequest true "User Login Request Body"
// @Success 200 {object} dto.LoginResponse "User Login Response Body"
// @Router /menu-planning-service/api/v1/login [post]
func Login(ctx *gin.Context) {
	var credential dto.LoginRequest
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	jwtWrapper := service.JwtWrapper{SecretKey: profile.GetStringValue("jwtConfig.secretKey"), Issuer: profile.GetStringValue("jwtConfig.issuer"), ExpirationHours: int64(profile.GetIntValue("jwtConfig.tokenExpiry"))}
	port := ports.UserLoginPort{UserLoginUseCase: &usecase.UserLoginUseCase{R: repository.GetUserRepository(postgres.GetDB())}}
	loginUser, err := port.LoginUser(credential.Email, credential.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Invalid email or password"})
		return
	}
	token, err := jwtWrapper.GenerateToken(loginUser.Email, loginUser.UserType)
	if err != nil {
		return
	}
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Invalid email or password"})
	}
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary User signup
// @Schemes
// @Description body
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param menuPlan body dto.UserSignUpRequest true "User SignUp Request Body"
// @Success 200 {object} dto.UserSignUpResponse "User SignUp Response Body"
// @Router /menu-planning-service/api/v1/signUp [post]
func SignUp(ctx *gin.Context) {
	var request dto.UserSignUpRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.UserSignUpPort{UserSignUpUseCase: &usecase.UserSignUpUseCase{R: repository.GetUserRepository(postgres.GetDB())}}
	ctx.JSON(http.StatusCreated, port.SignUp(&request))
}
