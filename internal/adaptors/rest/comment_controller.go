package rest

import (
	"github.com/gin-gonic/gin"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/ports"
	"hello-fresh-menu-planning-service/internal/application/usecase"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"net/http"
)

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Create Comment
// @Schemes
// @Description body
// @Tags Comment
// @Accept json
// @Produce json
// @Param comment body dto.CommentCreateRequest true "Comment create request body"
// @Success 200 {object} dto.CommentCreateResponse "Comment create response body"
// @Router /menu-planning-service/api/v1/comments [post]
func CreateComment(c *gin.Context) {
	var request dto.CommentCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.CommentCreatePort{CommentCreateUseCase: &usecase.CommentCreateUseCase{R: repository.GetCommentRepository(postgres.GetDB())}}
	response, err := port.CreateComment(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, &response)
}
