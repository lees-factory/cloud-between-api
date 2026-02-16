package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-domain/translation"
)

type TranslationController struct {
	service *translation.TranslationService
}

func NewTranslationController(service *translation.TranslationService) *TranslationController {
	return &TranslationController{service: service}
}

// GET /api/v1/translations?locale=ko
func (ctrl *TranslationController) GetAll(c *gin.Context) {
	locale := c.DefaultQuery("locale", "ko")

	result, err := ctrl.service.GetAll(c.Request.Context(), locale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET /api/v1/translations/:namespace?locale=ko
func (ctrl *TranslationController) GetByNamespace(c *gin.Context) {
	namespace := c.Param("namespace")
	locale := c.DefaultQuery("locale", "ko")

	result, err := ctrl.service.GetByNamespace(c.Request.Context(), locale, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "namespace not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
